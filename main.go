package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)

}

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		Scanner.Scan()
		userInputs := cleanInput(Scanner.Text())
		
		if len(userInputs) > 0{
		fmt.Printf("Your command was: %v\n", strings.ToLower(userInputs[0]))
		}
	}
}

