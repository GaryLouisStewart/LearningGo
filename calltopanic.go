package main

import "fmt"

func main() {

	demonstratePanic()
}

func demonstratePanic(){

	defer func() {

		fmt.Println(recover())
	}()

	panic("ABORT ABORT ABORT!")
}

