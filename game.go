package main

import (
	"fmt"
	"github.com/PyMarcus/go_gaming/source"
)


func main() {
	var name string
	
	fmt.Print("Player name: ")
	fmt.Scanf("%s", &name)

	source.PlayGame(name)
}
