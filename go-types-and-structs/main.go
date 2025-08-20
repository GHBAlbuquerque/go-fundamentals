package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{"Trevor","Go","+551189769747",10,time.Now()}

	user2 := User{
		FirstName:   "Grace",
		LastName:    "Go",
		PhoneNumber: "+551189769747",
		Age:         10,
		BirthDate:   time.Now(),
	}

	fmt.Println(user.FirstName)
	fmt.Println(user2.FirstName)
}

func ExportedFunc() {
	// accessible from other packages
}
