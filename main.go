package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliCommand struct{
	name string
	description string
	callback func() error
}

var SupportedCommands map[string]CliCommand =  make(map[string]CliCommand)


func Help() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println("\nhelp: Displays a help message")
		fmt.Println("exit: Exit the Pokedex\n")
		return nil
	}

func Exit() error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}


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




func cleanInput(text string) []string {
	return strings.Fields(text)

}

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

