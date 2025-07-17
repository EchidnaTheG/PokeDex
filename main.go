package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
	
}

func main() {
	value := cleanInput(" Hello World")
	fmt.Printf("%v\n%v\n",value[0],value[1])
}
