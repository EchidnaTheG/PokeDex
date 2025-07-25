package main

import (
	"bufio"
	"fmt"
	"github.com/EchidnaTheG/PokeDex/internal"
	"os"
	"strings"
)

// declaring CACHE type
var CACHE *internal.Cache

// General Struct For All CLI Commands In The App
type CliCommand struct {
	name        string
	description string
	callback    func(param string) error
}

// Initiliazing The Map That Contains All The Commands, the map is indexed using the names of the commands and the values are struct type CliCommand for that name
var SupportedCommands map[string]CliCommand = make(map[string]CliCommand)

// The Help CliCommand that is used by the CliCommand help in the SupportedCommands map as a callback method
func Help(_ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("\nhelp: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

// The exit CliCommand that is used by the CliCommand exit in the SupportedCommands map as a callback method
func Exit(_ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

// The map CliCommand that is used by the CliCommand map in the SupportedCommands map as a callback method, it used internal functions to manage the Api calls, not yet finalized
func Map(_ string) error {
	Locations, err := internal.GetLocationData(&config, false, CACHE)
	if err != nil {
		fmt.Printf("An Exception Has Ocurred: %v\n", err)
	}
	for _, Location := range Locations {
		fmt.Println(Location.Name)

	}
	return nil
}

func Mapb(_ string) error {

	Locations, err := internal.GetLocationData(&config, true, CACHE)
	if err != nil {
		fmt.Printf("An Exception Has Ocurred: %v\n", err)
	}
	for _, Location := range Locations {
		fmt.Println(Location.Name)

	}
	return nil
}

func Explore(area string) error {
	fmt.Printf("Exploring %v...\n", area)
	listOfPokemon, err := internal.GetPokemonInArea(area, CACHE)
	if err != nil {
		return err
	}
	fmt.Printf("Found Pokemon:\n")
	for _, data := range listOfPokemon {
		fmt.Printf(" - %v\n", data.Pokemon.Name)
	}

	return nil
}

func Catch(pokemon string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemon)
	SuccessBool, err := internal.TryStorePokemon(pokemon,&PokeDex, CACHE)
	if SuccessBool{
		fmt.Printf("%v was caught!\n", pokemon)
		return nil
	}
	if !SuccessBool &&  err == nil{
		fmt.Printf("%v escaped!\n", pokemon)
		return nil
	}

	return err
}
func Inspect(pokemon string) error{
	PokeDexPointer := &PokeDex
	var PokemonInPokedex internal.Pokemon
	validPokemon := false
	for _, PokemonInPokedex = range *PokeDexPointer{
		if PokemonInPokedex.Name == pokemon{
			validPokemon = true
			break
		}
	}
	if validPokemon{
		fmt.Printf("Name: %v\n",PokemonInPokedex.Name)
		fmt.Printf("Height: %v\n",PokemonInPokedex.Height)
		fmt.Printf("Weight: %v\n",PokemonInPokedex.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range PokemonInPokedex.Stats{
			fmt.Printf(" -%v: %v\n",stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, typ := range PokemonInPokedex.Types{
			fmt.Printf(" - %v\n",typ.Type.Name )
		}
		return nil
	}else{
		fmt.Printf("you have not caught that pokemon\n")
	}
	return nil	
}

func ShowPokedex(paramToCompleteSignature string) error{
	_ = paramToCompleteSignature
	pokedexPointer := &PokeDex
	fmt.Printf("Your Pokedex:\n")
	if len(*pokedexPointer) ==0{
		fmt.Printf("You Have No Pokemon!\n")
	}
	for _,pokem := range *pokedexPointer {
		fmt.Printf(" - %v\n",pokem.Name)
	}
	return nil
}


var config = internal.Config{Next: nil, Previous: nil}
var PokeDex []internal.Pokemon   //not using a map here because I want to support duplicate pokemon and add later pertinent features
//Initializing the SupportedCommands Map

func init() {
	CACHE = internal.NewCache(internal.INTERVAL)
	
	SupportedCommands = make(map[string]CliCommand)
	SupportedCommands["help"] = CliCommand{
		name:        "help",
		description: "Gives Help about app",
		callback:    Help,
	}
	SupportedCommands["exit"] = CliCommand{
		name:        "exit",
		description: "exits the program",
		callback:    Exit,
	}
	SupportedCommands["map"] = CliCommand{
		name:        "map",
		description: "lists all the locations in batches of 20",
		callback:    Map,
	}
	SupportedCommands["mapb"] = CliCommand{
		name:        "mapb",
		description: "lists all the locations in batches of 20, goes back 1 batch",
		callback:    Mapb,
	}
	SupportedCommands["mapb"] = CliCommand{
		name:        "mapb",
		description: "lists all the locations in batches of 20, goes back 1 batch",
		callback:    Mapb,
	}
	SupportedCommands["explore"] = CliCommand{
		name:        "explore",
		description: "lists all the pokemon in a specific area",
		callback:    Explore,
	}
	SupportedCommands["catch"] = CliCommand{
		name:        "catch",
		description: "attempts to catch the argument.",
		callback:    Catch,
	}
	SupportedCommands["inspect"] = CliCommand{
		name:        "inspect",
		description: "inspect a catched pokemon",
		callback:    Inspect,
	}
	SupportedCommands["pokedex"] = CliCommand{
		name:        "pokedex",
		description: "show all your catched pokemon",
		callback:    ShowPokedex,
	}
}

// Helper function for capturing user input
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))

}

// Main function and app loop. Scanner is initialized for input and a infinite loop starts, user input captured and text outputted, as well as commands called
func main() {
	Scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		Scanner.Scan()
		userInputs := cleanInput(Scanner.Text())

		if len(userInputs) > 0 {
			firstValue := userInputs[0]
			_, ok := SupportedCommands[firstValue]
			if !ok {
				println("Unknown Command")
				continue
			}
			switch firstValue {
			  case "explore", "catch", "inspect":
                if len(userInputs) <= 1 {
                    fmt.Println("Not Enough Flags Given")
                    continue
                }
                err := SupportedCommands[firstValue].callback(userInputs[1])
                if err != nil {
                    fmt.Printf("Error!: %v\n", err)
                }
			default:
				err := SupportedCommands[userInputs[0]].callback("")
				if err != nil {
					fmt.Printf("Error!: %v\n", err)
				}
			}

		}
	}
}
