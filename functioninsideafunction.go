package main

import "fmt"

func main() {


	num4 := 4

	doubleNum := func() int {

		num4 *= 2

		return num4
	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
}