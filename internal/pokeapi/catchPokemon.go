package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonData struct {
	id                     int
	Name                   string
	BaseExperience         int `json:"base_experience"`
	height                 int
	isDefault              bool
	order                  int
	weight                 int
	locationAreaEncounters string
	abilities              []struct {
		isHidden bool
		slot     int
		ability  struct {
			name string
			url  string
		}
	}
	forms []struct {
		name string
		url  string
	}
	gameIndices []struct {
		gameIndex int
		version   struct {
			name string
			url  string
		}
	}
	heldItems []struct {
		item struct {
			name string
			url  string
		}
		versionDetails []struct {
			rarity  int
			version struct {
				name string
				url  string
			}
		}
	}
	moves []struct {
		move struct {
			name string
			url  string
		}
		versionGroupDetails []struct {
			levelLearnedAt  int
			order           int
			moveLearnMethod struct {
				name string
				url  string
			}
		}
	}
	species struct {
		name string
		url  string
	}
	sprites struct {
		backDefault      *string
		backFemale       *string
		backShiny        *string
		backShinyFemale  *string
		frontDefault     *string
		frontFemale      *string
		frontShiny       *string
		frontShinyFemale *string
		other            struct {
			dreamWorld struct {
				frontDefault *string
				frontFemale  *string
			}
			home struct {
				frontDefault     *string
				frontFemale      *string
				frontShiny       *string
				frontShinyFemale *string
			}
			officialArtwork struct {
				frontDefault *string
				frontShiny   *string
			}
			showdown struct {
				backDefault      *string
				backFemale       *string
				backShiny        *string
				backShinyFemale  *string
				frontDefault     *string
				frontFemale      *string
				frontShiny       *string
				frontShinyFemale *string
			}
		}
		versions struct {
			generationI struct {
				redBlue struct {
					backDefault  *string
					backGray     *string
					frontDefault *string
					frontGray    *string
				}
				yellow struct {
					backDefault  *string
					backGray     *string
					frontDefault *string
					frontGray    *string
				}
			}
			generationII struct {
				crystal struct {
					backDefault  *string
					backShiny    *string
					frontDefault *string
					frontShiny   *string
				}
				gold struct {
					backDefault  *string
					backShiny    *string
					frontDefault *string
					frontShiny   *string
				}
				silver struct {
					backDefault  *string
					backShiny    *string
					frontDefault *string
					frontShiny   *string
				}
			}
			generationIII struct {
				emerald struct {
					frontDefault *string
					frontShiny   *string
				}
				fireredLeafgreen struct {
					backDefault  *string
					backShiny    *string
					frontDefault *string
					frontShiny   *string
				}
				rubySapphire struct {
					backDefault  *string
					backShiny    *string
					frontDefault *string
					frontShiny   *string
				}
			}
			generationIV struct {
				diamondPearl struct {
					backDefault      *string
					backFemale       *string
					backShiny        *string
					backShinyFemale  *string
					frontDefault     *string
					frontFemale      *string
					frontShiny       *string
					frontShinyFemale *string
				}
				heartgoldSoulsilver struct {
					backDefault      *string
					backFemale       *string
					backShiny        *string
					backShinyFemale  *string
					frontDefault     *string
					frontFemale      *string
					frontShiny       *string
					frontShinyFemale *string
				}
				platinum struct {
					backDefault      *string
					backFemale       *string
					backShiny        *string
					backShinyFemale  *string
					frontDefault     *string
					frontFemale      *string
					frontShiny       *string
					frontShinyFemale *string
				}
			}
			generationV struct {
				blackWhite struct {
					backDefault      *string
					backFemale       *string
					backShiny        *string
					backShinyFemale  *string
					frontDefault     *string
					frontFemale      *string
					frontShiny       *string
					frontShinyFemale *string
					animated         struct {
						backDefault      *string
						backFemale       *string
						backShiny        *string
						backShinyFemale  *string
						frontDefault     *string
						frontFemale      *string
						frontShiny       *string
						frontShinyFemale *string
					}
				}
			}
			generationVI struct {
				omegarubyAlphasaphire struct {
					frontDefault     *string
					frontFemale      *string
					frontShiny       *string
					frontShinyFemale *string
				}
				xY struct {
					frontDefault     *string
					frontFemale      *string
					frontShiny       *string
					frontShinyFemale *string
				}
			}
			generationVII struct {
				icons struct {
					frontDefault *string
					frontFemale  *string
				}
				ultrasunUltramoon struct {
					frontDefault     *string
					frontFemale      *string
					frontShiny       *string
					frontShinyFemale *string
				}
			}
			generationVIII struct {
				icons struct {
					frontDefault *string
					frontFemale  *string
				}
			}
		}
	}
	cries struct {
		latest string
		legacy string
	}
	stats []struct {
		baseStat int
		effort   int
		stat     struct {
			name string
			url  string
		}
	}
	types []struct {
		slot    int
		theType struct {
			name string
			url  string
		}
	}
	pastTypes []struct {
		generation struct {
			name string
			url  string
		}
		types []struct {
			slot    int
			theType struct {
				name string
				url  string
			}
		}
	}
	pastAbilities []struct {
		generation struct {
			name string
			url  string
		}
		abilities []struct {
			ability  *string
			isHidden bool
			slot     int
		}
	}
}

func (c *PokeClient) GetPokemon(pokemonName string) (PokemonData, error) {
	endpoint := baseApi + "pokemon/" + pokemonName

	if data, ok := c.cache.Get(endpoint); ok {
		var jsonData PokemonData
		err := json.Unmarshal(data, &jsonData)
		if err != nil {
			return PokemonData{}, err
		}
		return jsonData, nil
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return PokemonData{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return PokemonData{}, fmt.Errorf("Server sent bad code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonData{}, err
	}

	var jsonData PokemonData
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return PokemonData{}, err
	}
	return jsonData, nil
}
