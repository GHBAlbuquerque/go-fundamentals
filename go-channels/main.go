package main

import (
	"fmt"
	"go-channels/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	n := helpers.RandomNumber(numPool)
	intChan <- n
}

func main() {
	intChan := make(chan int) // channels are used to send information that will be received in one of more places of the program, asynchronously
	defer close(intChan)      // important! all channels need to be closed. With defer, they close after program execution

	go CalculateValue(intChan)

	int := <-intChan // this is how you get the value out of the channel after exec
	fmt.Println(int)
}
