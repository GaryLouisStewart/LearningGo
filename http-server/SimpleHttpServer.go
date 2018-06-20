package main

import (
	"net/http"
	"fmt"
)
func main() {

	http.HandleFunc("/", handler)

	http.HandleFunc("/remind", handler2)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Give me stars on github\n")
}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Remember to give me stars on github")
}

