package main

import "fmt"

func main() {

	//variables below messing around setting vars and printing them out....

	var firstName string = "Gary"
	var middleName string = "Louis"
	var surName string = "Stewart"
	var age int = 42
	var favNum float64 = 1.1278193
	var favDrink string = "Macallan Neat"

	//fmt.Println(firstName)
	//fmt.Println(middleName)
	//fmt.Println(surName)
	//fmt.Println(age)
	//fmt.Println(favNum)
	//fmt.Println(favDrink)
	fmt.Println(firstName, middleName, surName, age, favNum, favDrink)

	// arithmetic operation examples

	var a = 6 * 8
	var b = 6 - 8
	var c = 6 + 8
	var d = 6 / 8
	var e = 6 % 8
	//fmt.Println("6 * 8 =", 6 * 8)
	//fmt.Println("6 - 8 =", 6 - 8)
	//fmt.Println("6 + 8 =", 6 + 8)
	//fmt.Println("6 / 8 =", 6 / 8)
	//fmt.Println("6 % 8 =", 6 % 8)
	fmt.Println(a, b, c, d, e)
}
