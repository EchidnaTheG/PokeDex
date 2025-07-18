package internal

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// PokeApi
const API string = "https://pokeapi.co/api/v2/"

//struct for location api call
type locationData struct{
	ID int `json:"id"`
	Name string `json:"name"`
}



//function for getting API location data
func GetLocationData() (locationData, error){
	FullPath := API + "location"
	res, err := http.Get(FullPath)
	if err != nil{
		return locationData{}, fmt.Errorf("connection Error: %v", err)
	}
	var NewLocationData locationData
	Decoder := json.NewDecoder(res.Body)
	err = Decoder.Decode(&NewLocationData)
	if err != nil{
		return locationData{}, fmt.Errorf("error Decoding: %v", err)
	}
	return NewLocationData, nil
}