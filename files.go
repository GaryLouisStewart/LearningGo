package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
)

func main() {
	file, err := os.Create("sample.txt")

	if err != nil {

		log.Fatal(err)

	}

	file.WriteString("this is educational")

	file.Close()

	stream, err := ioutil.ReadFile("sample.txt")

	if err != nil {

		log.Fatal(err)
	}

	readString := string(stream)

	fmt.Println(readString)
}