package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

// PokeApi
const API string = "https://pokeapi.co/api/v2"

// function for getting API location data
func GetLocationData(configPointer *Config, isMapb bool, CACHE *Cache) (Results []LocationResult, err error) {
	var FullPath string

	if configPointer.Next == nil && configPointer.Previous == nil {
		FullPath = API + "/location-area"
	}

	if configPointer.Next != nil {
		FullPath = *configPointer.Next
	}

	if isMapb {
		if configPointer.Previous == nil {
			fmt.Printf("Cannot Use Mapb\n")
			return []LocationResult{}, nil
		}
		FullPath = *configPointer.Previous
	}

	if value, ok := CACHE.Get(FullPath); ok {

		var NewLocationData LocationData
		Decoder := json.NewDecoder(bytes.NewReader(value))
		err = Decoder.Decode(&NewLocationData)
		if err != nil {
			return []LocationResult{}, fmt.Errorf("error Decoding Json: %v", err)
		}

		configPointer.Next = NewLocationData.Next
		configPointer.Previous = NewLocationData.Previous
		var LocationResult []LocationResult
		LocationResult = append(LocationResult, NewLocationData.Results...)
		return LocationResult, nil

	}

	res, err := http.Get(FullPath)
	if err != nil {
		return []LocationResult{}, fmt.Errorf("connection Error: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return []LocationResult{}, fmt.Errorf("bad Request: %v", res.StatusCode)
	}
	defer res.Body.Close()
	cacheValue, err := io.ReadAll(res.Body)
	if err != nil {
		return []LocationResult{}, fmt.Errorf("error reading response body: %v", err)
	}

	var NewLocationData LocationData
	Decoder := json.NewDecoder(bytes.NewReader(cacheValue))
	err = Decoder.Decode(&NewLocationData)
	if err != nil {
		return []LocationResult{}, fmt.Errorf("error Decoding Json: %v", err)
	}
	configPointer.Next = NewLocationData.Next
	configPointer.Previous = NewLocationData.Previous

	CACHE.Add(FullPath, cacheValue)
	var LocationResult []LocationResult
	LocationResult = append(LocationResult, NewLocationData.Results...)
	return LocationResult, nil
}

func GetPokemonInArea(area string, CACHE *Cache) (Pokemon []PokemonEncounter, err error) {

	FullPath := API + "/location-area/" + area

	if value, ok := CACHE.Get(FullPath); ok {

		var LocationAreaInfo LocationAreaData
		Decoder := json.NewDecoder(bytes.NewReader(value))
		err = Decoder.Decode(&LocationAreaInfo)
		if err != nil {
			return []PokemonEncounter{}, fmt.Errorf("error Decoding Json: %v", err)
		}
		return LocationAreaInfo.PokemonEncounters, nil

	}

	res, err := http.Get(FullPath)
	if err != nil {
		return []PokemonEncounter{}, err
	}
	if res.StatusCode != http.StatusOK {
		return []PokemonEncounter{}, fmt.Errorf("bad Request: %v", res.StatusCode)
	}

	defer res.Body.Close()
	cacheValue, err := io.ReadAll(res.Body)
	if err != nil {
		return []PokemonEncounter{}, fmt.Errorf("error reading response body: %v", err)
	}

	var LocationAreaInfo LocationAreaData
	Decoder := json.NewDecoder(bytes.NewReader(cacheValue))
	err = Decoder.Decode(&LocationAreaInfo)
	if err != nil {
		return []PokemonEncounter{}, fmt.Errorf("error Decoding Json: %v", err)
	}
	CACHE.Add(FullPath, cacheValue)
	return LocationAreaInfo.PokemonEncounters, nil

}

func CatchPokemonSuccess(PokemonBaseExperience int) (outcome bool) {
	r := rand.Float64()

	switch {
	case PokemonBaseExperience <= 85:
		return r >= 0.1
	case PokemonBaseExperience <= 110:
		return r >= 0.2
	case PokemonBaseExperience <= 130:
		return r >= 0.25
	case PokemonBaseExperience <= 150:
		return r >= 0.35
	case PokemonBaseExperience <= 190:
		return r >= 0.4
	case PokemonBaseExperience <= 220:
		return r >= 0.5
	case PokemonBaseExperience <= 250:
		return r >= 0.6
	case PokemonBaseExperience <= 280:
		return r >= 0.7
	case PokemonBaseExperience <= 300:
		return r >= 0.8
	case PokemonBaseExperience <= 350:
		return r >= 0.9
	default:
		return r >= 0.95 // Anything > 350
	}
}
func TryStorePokemon(pokemon string, PokeDex *[]Pokemon, CACHE *Cache) (bool, error) {
	FullPath := API + "/pokemon/" + strings.ToLower(pokemon)
	if value, ok := CACHE.Get(FullPath); ok {

		var PokemonInfo Pokemon
		Decoder := json.NewDecoder(bytes.NewReader(value))
		err := Decoder.Decode(&PokemonInfo)
		if err != nil {
			return false, fmt.Errorf("error Decoding Json: %v", err)
		}
		if CatchPokemonSuccess(PokemonInfo.BaseExperience){
			*PokeDex = append(*PokeDex,PokemonInfo)
			return true, nil
		}
		return false, nil

	}

	res, err := http.Get(FullPath)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("bad Request: %v", res.StatusCode)
	}
	cacheValue, err := io.ReadAll(res.Body)
	if err != nil {
		return false, fmt.Errorf("error reading response body: %v", err)
	}
	var PokemonInfo Pokemon
	Decoder := json.NewDecoder(bytes.NewReader(cacheValue))
	err = Decoder.Decode(&PokemonInfo)
	if err != nil {
		return false, fmt.Errorf("error Decoding Json: %v", err)
	}
	CACHE.Add(FullPath, cacheValue)
	if CatchPokemonSuccess(PokemonInfo.BaseExperience){
			*PokeDex = append(*PokeDex,PokemonInfo)
			return true, nil
		}
	return false, nil
}
