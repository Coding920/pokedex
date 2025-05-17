package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ExploreJson struct {
	id                   int
	name                 string
	gameIndex            int
	encounterMethodRates []struct {
		encounterMethod struct {
			name string
			url  string
		}
		versionDetails []struct {
			rate    int
			version struct {
				name string
				url  string
			}
		}
	}
	location struct {
		name string
		url  string
	}
	names []struct {
		name     string
		language struct {
			name string
			url  string
		}
	}
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			url  string
		} `json:"pokemon"`
		versionDetails []struct {
			version struct {
				name string
				url  string
			}
			maxChance        int
			encounterDetails []struct {
				minLevel        int
				maxLevel        int
				conditionValues []any
				chance          int
				method          struct {
					name string
					url  string
				}
			}
		}
	} `json:"pokemon_encounters"`
}

func (c *PokeClient) ExploreLocation(location string) (ExploreJson, error) {
	endpoint := baseApi + "location-area/" + location

	if data, ok := c.cache.Get(endpoint); ok {
		var jsonData ExploreJson
		err := json.Unmarshal(data, &jsonData)
		if err != nil {
			return ExploreJson{}, err
		}
		return jsonData, nil
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return ExploreJson{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return ExploreJson{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return ExploreJson{}, fmt.Errorf("Server response not OK, response: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreJson{}, err
	}
	c.cache.Add(endpoint, data)

	var jsonData ExploreJson
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Err unmarshalling")
		fmt.Println(err)
		return ExploreJson{}, err
	}

	return jsonData, nil
}
