package main

import (
	"strconv"
	"fmt"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string){
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough and Send for Sauce");

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)

}

func makeSauce(stringChan chan string){
	pizza := <- stringChan

	fmt.Println("Add Sauce and Send", pizza, "for toppings");

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)

}

func addToppings(stringChan chan string){

	pizza := <- stringChan

	fmt.Println("Add Toppings to", pizza, "and ship")

	time.Sleep(time.Millisecond * 10)

}

func main() {

	stringChan := make(chan string)

	for i := 0; i < 3; i++{

		go makeDough(stringChan)
		go makeSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}

}