package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *PokeClient) GetLocations(url *string) (LocationsJson, error) {
	endpoint := baseApi + "location-area/"
	if url != nil {
		endpoint = *url
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
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&jsonData)
	if err != nil {
		return LocationsJson{}, err
	}
	return jsonData, nil
}
