package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/EchidnaTheG/PokeDex/internal"
)



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
		fmt.Println("exit: Exit the Pokedex\n")
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
	internal.GetLocationData()
}

//Initializing the SupportedCommands Map
func init(){
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

