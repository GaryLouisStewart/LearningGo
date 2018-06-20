package main

import "fmt"

func main() {

	y := 0

	changeYValNow(&y)

	fmt.Println("y =", y)

	fmt.Println("Memory Address for y =", &y)
}

func changeYValNow(y *int){

	*y = 2

}