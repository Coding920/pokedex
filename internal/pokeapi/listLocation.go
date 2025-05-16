package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *PokeClient) GetLocations(url *string) (LocationsJson, error) {
	endpoint := baseApi + "location-area/"
	if url != nil {
		endpoint = *url
	}
	data, ok := c.cache.Get(endpoint)
	if ok {
		var jsonData LocationsJson
		err := json.Unmarshal(data, &jsonData)
		if err != nil {
			return LocationsJson{}, err
		}

		return jsonData, nil
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return LocationsJson{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return LocationsJson{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return LocationsJson{},
			fmt.Errorf("Bad Server response, Code: %v", res.StatusCode)
	}

	var jsonData LocationsJson
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsJson{}, err
	}

	c.cache.Add(endpoint, resBody)
	err = json.Unmarshal(resBody, &jsonData)
	if err != nil {
		return LocationsJson{}, err
	}

	return jsonData, nil
}
