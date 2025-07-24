package internal

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"io"

)

// PokeApi
const API string = "https://pokeapi.co/api/v2"


type LocationAreaData struct {
    ID                   int                    `json:"id"`
    Name                 string                 `json:"name"`
    GameIndex            int                    `json:"game_index"`
    EncounterMethodRates []EncounterMethodRate  `json:"encounter_method_rates"`
    Location             LocationResult                `json:"location"`
    Names                []AreaName             `json:"names"`
    PokemonEncounters    []PokemonEncounter     `json:"pokemon_encounters"`
}

// Encounter method rate structure
type EncounterMethodRate struct {
    EncounterMethod LocationResult               `json:"encounter_method"`
    VersionDetails  []MethodVersionDetail `json:"version_details"`
}

// Version details for encounter methods
type MethodVersionDetail struct {
    Rate    int     `json:"rate"`
    Version LocationResult `json:"version"`
}

// Area name structure
type AreaName struct {
    Name     string  `json:"name"`
    Language LocationResult `json:"language"`
}

// Pokemon encounter structure
type PokemonEncounter struct {
    Pokemon        LocationResult                      `json:"pokemon"`
    VersionDetails []PokemonEncounterVersionDetail `json:"version_details"`
}

// Version details for pokemon encounters
type PokemonEncounterVersionDetail struct {
    Version          LocationResult            `json:"version"`
    MaxChance        int                `json:"max_chance"`
    EncounterDetails []EncounterDetail  `json:"encounter_details"`
}

// Individual encounter detail
type EncounterDetail struct {
    MinLevel        int       `json:"min_level"`
    MaxLevel        int       `json:"max_level"`
    ConditionValues []LocationResult `json:"condition_values"`
    Chance          int       `json:"chance"`
    Method          LocationResult   `json:"method"`
}

// Individual location result
type LocationResult struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

// Main response structure
type LocationData struct {
    Count    int              `json:"count"`
    Next     *string          `json:"next"`     // pointer because it can be null
    Previous *string          `json:"previous"` // pointer because it can be null
    Results  []LocationResult `json:"results"`  // array of location results
}

type Config struct{
	Next *string
	Previous *string
}

//function for getting API location data
func GetLocationData(configPointer *Config, isMapb bool, CACHE *Cache ) (Results []LocationResult,err error){
	var FullPath string

	if (configPointer.Next == nil && configPointer.Previous == nil){
		FullPath = API + "/location-area"
	}

	if configPointer.Next != nil{
		FullPath= *configPointer.Next 
	}
	

	if isMapb {
		if configPointer.Previous == nil{
			fmt.Printf("Cannot Use Mapb\n")
			return []LocationResult{}, nil
		}
		FullPath= *configPointer.Previous
	}

	
	if value,ok := CACHE.Get(FullPath); ok{

		var NewLocationData LocationData
		Decoder := json.NewDecoder(bytes.NewReader(value))
		err = Decoder.Decode(&NewLocationData)
		if err != nil{
			return []LocationResult{}, fmt.Errorf("error Decoding Json: %v", err)
		}

		configPointer.Next= NewLocationData.Next
		configPointer.Previous= NewLocationData.Previous
		var LocationResult []LocationResult
		LocationResult = append(LocationResult, NewLocationData.Results...)
		return LocationResult, nil

	}
	
	
	res, err := http.Get(FullPath)
	if err != nil{
		return []LocationResult{}, fmt.Errorf("connection Error: %v", err)
	}
	if res.StatusCode != http.StatusOK{
		return []LocationResult{}, fmt.Errorf("Bad Request: %v\n", res.StatusCode)
	}
	defer res.Body.Close()
	cacheValue, err := io.ReadAll(res.Body)
    if err != nil {
        return []LocationResult{}, fmt.Errorf("error reading response body: %v", err)
    }
	
	var NewLocationData LocationData
	Decoder := json.NewDecoder(bytes.NewReader(cacheValue))
	err = Decoder.Decode(&NewLocationData)
	if err != nil{
		return []LocationResult{}, fmt.Errorf("error Decoding Json: %v", err)
	}
	configPointer.Next= NewLocationData.Next
	configPointer.Previous= NewLocationData.Previous
	
	CACHE.Add(FullPath,cacheValue)
	var LocationResult []LocationResult
	LocationResult = append(LocationResult, NewLocationData.Results...)
	return LocationResult, nil
}

func GetPokemonInArea(area string, CACHE *Cache) (Pokemon []PokemonEncounter, err error){
	
	FullPath := API + "/location-area/" + area


	if value,ok := CACHE.Get(FullPath); ok{

		var LocationAreaInfo LocationAreaData
		Decoder := json.NewDecoder(bytes.NewReader(value))
		err = Decoder.Decode(&LocationAreaInfo)
		if err != nil{
			return []PokemonEncounter{}, fmt.Errorf("error Decoding Json: %v", err)
		}
		return LocationAreaInfo.PokemonEncounters, nil

		
	}


	res, err := http.Get(FullPath)
	if err != nil{
		return []PokemonEncounter{}, err
	}
	if res.StatusCode != http.StatusOK{
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
	if err != nil{
		return []PokemonEncounter{}, fmt.Errorf("error Decoding Json: %v", err)
	}
	CACHE.Add(FullPath,cacheValue)
	return LocationAreaInfo.PokemonEncounters, nil

}