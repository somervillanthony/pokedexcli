package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UnmarshalJson(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("HTTP GET failed: %w", err)
	}
	defer res.Body.Close()

	var jsonData LocationAreas
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&jsonData)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("failed to decode json: %w", err)
	}

	return jsonData, nil
}
