package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
type AreaDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

var cache *Cache = NewCache(5)

func GetLocationArea(link string) LocationArea {
	val, ok := cache.Get(link)
	if ok {
		locArea := LocationArea{}
		err := json.Unmarshal(val, &locArea)
		if err != nil {
			log.Fatal(err)
		}
		return locArea
	}
	res, err := http.Get(link)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	locArea := LocationArea{}
	err = json.Unmarshal(body, &locArea)
	if err != nil {
		log.Fatal(err)
	}
	cache.Add(link, body)

	return locArea
}

const DEFAULT_LOC_AREA string = "https://pokeapi.co/api/v2/location-area/"

func GetPokemonsFromLocArea(areaName string) AreaDetails {
	val, ok := cache.Get(DEFAULT_LOC_AREA + areaName)
	fmt.Println(DEFAULT_LOC_AREA + areaName)
	if ok {
		areaDetails := AreaDetails{}
		err := json.Unmarshal(val, &areaDetails)
		if err != nil {
			log.Fatal(err)
		}
		return areaDetails
	}
	res, err := http.Get(DEFAULT_LOC_AREA + areaName)

	if err != nil {
		log.Fatal(err)

	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	areaDetails := AreaDetails{}
	err = json.Unmarshal(body, &areaDetails)
	if err != nil {
		log.Fatal(err)
	}
	cache.Add(DEFAULT_LOC_AREA+areaName, body)

	return areaDetails
}
