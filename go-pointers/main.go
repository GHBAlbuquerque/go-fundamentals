package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString) // myString is set to Green

	changeUsingPointer(&myString)

	log.Println("myString is now set to", myString) //myString is now set to Red
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s) //s is set to 0x14000118030
	newValue := "Red"

	*s = newValue
}
