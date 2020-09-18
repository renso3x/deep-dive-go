package main

import "fmt"

func main() {
	// Initialize maps
	colors := map[string]string {
		"red": "#ff0000",
		"blue": "#fefefe",
	}

	fmt.Println(colors)

	// Initialize maps
	car := make(map[string]string)
	// zero value
	fmt.Println(car)

	car["toyota"] = "wigo"
	car["10"] = "100000"

	printMap(car)
	// delete(car, "10")
	// fmt.Println(car)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("key " + k + ", value: " + v)
	}
}