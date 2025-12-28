package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) PokemonInfo(pokemonName string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// first check cache
	if data, ok := c.cache.Get(url); ok {
		var result PokemonResponse

		if err := json.Unmarshal(data, &result); err != nil {
			return PokemonResponse{}, fmt.Errorf("error unmarshalling pokemon response data: %w", err)
		}

		return result, nil
	}

	// ...otherwise just make the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return PokemonResponse{}, fmt.Errorf("error doing the GET request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonResponse{}, fmt.Errorf("error reading result body: %w", err)
	}

	c.cache.Add(url, data)

	var result PokemonResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return PokemonResponse{}, fmt.Errorf("error unmarshaling the pokemon info: %w", err)
	}

	return result, nil
}
