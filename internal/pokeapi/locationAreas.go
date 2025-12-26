package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) GetLocationAreas(url string) (LocationAreasResponse, error) {
	// check cache first
	if val, ok := c.cache.Get(url); ok {
		var result LocationAreasResponse
		if err := json.Unmarshal(val, &result); err != nil {
			return LocationAreasResponse{}, fmt.Errorf("error while unmarshalling: %w", err)
		}

		// fmt.Println("Value was received from cache !!!!")
		return result, nil
	}
	// fmt.Println("VALUE WAS NOT RECEIVED FROM CACHE DUBBMADSSSSSS!!!121")

	// create a new api call if not in cache
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error creating GET request: %w", err)
	}

	res, err := c.Client.Do(req)
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

	c.cache.Add(url, data)

	var result LocationAreasResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error while unmarshalling: %w", err)
	}

	return result, nil
}
