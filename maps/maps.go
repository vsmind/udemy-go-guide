package main

import (
	"fmt"
)

func main() {
	// Declare map
	var colors1 map[string]string = map[string]string{}

	colors2 := make(map[string]string)

	colors3 := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	// Add values to map

	colors1["white"] = "#FFFFFF"

	fmt.Println(colors1)
	fmt.Println(colors2)
	fmt.Println(colors3)

	// Remove key/value pair from map with build in funcion
	delete(colors1, "white")

	fmt.Println(colors1)

	// Iterate over map
	printMap(colors3)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color", color, "hex", hex)
	}
}
