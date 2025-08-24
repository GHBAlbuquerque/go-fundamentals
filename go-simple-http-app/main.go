package main

import (
	"errors"
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
	result := addValues(2, 2)
	text := fmt.Sprintf("This is the about page. Sum is: %d", result)
	_, _ = fmt.Fprintf(writer, text)
}

// Divide is the about page handler
func Divide(writer http.ResponseWriter, r *http.Request) {
	result, err := divideValues(100, 0)

	if err != nil {
		fmt.Fprintf(writer, err.Error())
		return
	}

	text := fmt.Sprintf("This is the divide page. Division is: %f", result)
	_, _ = fmt.Fprintf(writer, text)
}

// addValues adds two integers and returns the sum
func addValues(x, y int) int {
	return x + y
}

// divideValues adds two integers and returns the sum
func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

func main() {
	//built in package of the standard library
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf(fmt.Sprintf("Starting server on port %s...\n", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
