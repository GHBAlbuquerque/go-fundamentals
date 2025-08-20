package main

import "fmt"

var s = "text"

func main() {
	s := 8
	fmt.Println(s)

	fmt.Println("Hello World")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, cruel world!"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i prints to", i)

	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)

	anotherThingSaid, yetAnotherThingSaid := saySomethings()

	fmt.Println("The function returned", anotherThingSaid, "and", yetAnotherThingSaid)

	var g = 8
	h := 9

	fmt.Println(g, h)

	sayS("new text")
}

func saySomething() string {
	return "something"
}

func saySomethings() (string, string) {
	return "something", "another something"
}

func sayS(s string) {
	println(s)
}
