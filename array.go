package main

import "fmt"

func main() {

	var favNums [5] float64
	favNums[0] = 163
	favNums[1] = 78557
	favNums[2] = 691
	favNums[3] = 3.141
	favNums[4] = 1.618

	fmt.Println(favNums[3])

	favnums3 := [5]float64{1, 2, 3, 4, 5}

	for i, value := range favnums3 {
		fmt.Println(value, i)
	}
}