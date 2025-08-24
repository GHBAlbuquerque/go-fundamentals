package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "This is the home page")
}

// About is the about page handler
func About(writer http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	text := fmt.Sprintf("This is the about page. Sum is: %d", sum)
	_, _ = fmt.Fprintf(writer, text)
}

// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

func main() {
	//built in package of the standard library
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf(fmt.Sprintf("Starting server on port %s...\n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
