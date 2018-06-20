package main

import "fmt"

func main() {

	defer secondPrint()
	firstPrint()

}

func firstPrint(){ fmt.Println(1)}
func secondPrint(){ fmt.Println(2)}