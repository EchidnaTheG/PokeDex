package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/EchidnaTheG/PokeDex/internal"

)

//declaring CACHE type early
var CACHE *internal.Cache

// General Struct For All CLI Commands In The App
type CliCommand struct{
	name string
	description string
	callback func() error
}

// Initiliazing The Map That Contains All The Commands, the map is indexed using the names of the commands and the values are struct type CliCommand for that name
var SupportedCommands map[string]CliCommand =  make(map[string]CliCommand)



// The Help CliCommand that is used by the CliCommand help in the SupportedCommands map as a callback method
func Help() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println("\nhelp: Displays a help message")
		fmt.Println("exit: Exit the Pokedex")
		return nil
	}

// The exit CliCommand that is used by the CliCommand exit in the SupportedCommands map as a callback method	
func Exit() error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}
// The map CliCommand that is used by the CliCommand map in the SupportedCommands map as a callback method, it used internal functions to manage the Api calls, not yet finalized	
func Map() error{
	Locations, err := internal.GetLocationData(&config,false,CACHE)
	if err != nil{
		fmt.Printf("An Exception Has Ocurred: %v\n",err)
	}
	for _, Location := range Locations{
		fmt.Println(Location.Name)
	
	}
	return nil
}

func Mapb() error{
	Locations, err := internal.GetLocationData(&config,true,CACHE)
	if err != nil{
		fmt.Printf("An Exception Has Ocurred: %v\n",err)
	}
	for _, Location := range Locations{
		fmt.Println(Location.Name)
	
	}
	return nil
}
var config= internal.Config{Next:nil,Previous:nil}
//Initializing the SupportedCommands Map

func init(){
	CACHE = internal.NewCache(internal.INTERVAL)

	SupportedCommands=  make(map[string]CliCommand)
	SupportedCommands["help"] =CliCommand{
	name: "help",
	description: "Gives Help about app",
	callback: Help,
	}
	SupportedCommands["exit"]=CliCommand{
		name: "exit",
		description:"exits the program",
		callback: Exit,
	}
	SupportedCommands["map"]= CliCommand{
		name:"map",
		description: "lists all the locations in batches of 20",
		callback: Map,
	}
	SupportedCommands["mapb"]= CliCommand{
		name:"mapb",
		description: "lists all the locations in batches of 20, goes back 1 batch",
		callback: Mapb,
	}

	
} 



// Helper function for capturing user input
func cleanInput(text string) []string {
	return strings.Fields(text)

}


// Main function and app loop. Scanner is initialized for input and a infinite loop starts, user input captured and text outputted, as well as commands called
func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		Scanner.Scan()
		userInputs := cleanInput(strings.ToLower(Scanner.Text()))
		
		if len(userInputs) > 0{
			_,ok := SupportedCommands[userInputs[0]]; if !ok{
				println("Unknown Command")
				continue
			}
			err := SupportedCommands[userInputs[0]].callback()
			if err != nil{
				fmt.Printf("Error!: %v\n", err)
			}
		}
	}
}

