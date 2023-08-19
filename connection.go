package pokeapi

import (
	"encoding/json"
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

func GetLocationArea(link string) LocationArea {
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

	return locArea
}
