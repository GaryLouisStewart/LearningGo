package main

import "fmt"

func main() {

	fmt.Println(safeDivide(3,0))
	fmt.Println(safeDivide(3,2))

}

func safeDivide(num1, num2 int) int {

	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}