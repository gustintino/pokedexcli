package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreasResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreasResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error while doing a GET request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %d %s", res.StatusCode, res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error while reading result body: %w", err)
	}

	var result LocationAreasResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error while unmarshalling: %w", err)
	}

	return result, nil
}
