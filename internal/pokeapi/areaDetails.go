package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// NOTE: THIS IS NOT FINISHED
func (c Client) AreaDetails(areaName string) (AreaDetailsResponse, error) {
	url := baseURL + "/location-area/" + areaName

	// on standby, ready to fire sir
	// fmt.Printf("[DEBUG]: url being used: %s", url)

	// 1. First check if area details are in cache
	if data, ok := c.cache.Get(url); ok {
		var result AreaDetailsResponse
		if err := json.Unmarshal(data, &result); err != nil {
			return AreaDetailsResponse{}, fmt.Errorf("error unmarshalling AreaDetails data: %w", err)
		}

		fmt.Println("[DEUBG]: This data was gotten from the cache.")
		return result, nil
	}

	// NOTE: I don't have to check the cache for if the link exists, as I can easily build it myself.
	// It's probably quite a bit more expensive to do it anyway,
	// as I would need to go through each cache entry and check every location.

	// 2. If not we just make the link and send a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaDetailsResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return AreaDetailsResponse{}, fmt.Errorf("error doing the GET request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return AreaDetailsResponse{}, fmt.Errorf("bad status code: %d %s", res.StatusCode, res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return AreaDetailsResponse{}, fmt.Errorf("error reading result body: %w", err)
	}

	c.cache.Add(url, data)

	var result AreaDetailsResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return AreaDetailsResponse{}, fmt.Errorf("error unmarshaling area details: %w", err)
	}

	return result, nil
}
