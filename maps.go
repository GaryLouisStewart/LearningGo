package main

import "fmt"

func main() {
	presidentAge := make(map[string] int)

	presidentAge["TheodoreRoosevelt"] = 42

	fmt.Println(len(presidentAge))

	presidentAge["John F. Kennedy"] = 43
	fmt.Println(len(presidentAge))

	delete(presidentAge, "John F. Kennedy")
	fmt.Println(len(presidentAge))
}
