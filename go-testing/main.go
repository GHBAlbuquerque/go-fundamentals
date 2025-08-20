package main

import (
	"errors"
	"log"
)

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot divide by zero")
	}

	result = x / y
	return result, nil
}

func main() {
	result, err := divide(100.0, 20.0)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Result is", result)
}
