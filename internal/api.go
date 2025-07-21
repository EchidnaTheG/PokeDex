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