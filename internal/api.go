package internal

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// PokeApi
const API string = "https://pokeapi.co/api/v2"


// Individual location result
type LocationResult struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

// Main response structure
type LocationData struct {
    Count    int              `json:"count"`
    Next     string          `json:"next"`     // pointer because it can be null
    Previous string          `json:"previous"` // pointer because it can be null
    Results  []LocationResult `json:"results"`  // array of location results
}

type Config struct{
	Next string
	Previous string
}

//function for getting API location data
func GetLocationData(configPointer *Config, isMapb bool) (Results []LocationResult,err error){
	var FullPath string
	if (configPointer.Next == "" && configPointer.Previous == ""){
		FullPath = API + "/location-area"
	}else if isMapb{
		FullPath= configPointer.Previous
	}else{
		FullPath= configPointer.Next 
	}

	
	
	res, err := http.Get(FullPath)
	if err != nil{
		return []LocationResult{}, fmt.Errorf("connection Error: %v", err)
	}
	var NewLocationData LocationData
	Decoder := json.NewDecoder(res.Body)
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