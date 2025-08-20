package main

import (
	"fmt"
	"go-structs-with-functions/mypackage"
)

type myStruct struct {
	FirstName string
}

func (m *myStruct) getFirstName() string {
	return m.FirstName
}

func (m *myStruct) changeFirstName(name string) {
	m.FirstName = name
}

func (m myStruct) changeCopyFirstName(name string) {
	m.FirstName = name
}

func main() {
	// 1. Explicit struct creation
	var myVar myStruct
	myVar.FirstName = "John"

	fmt.Println(myVar)
	fmt.Println(myVar.getFirstName())

	// 2. Explicit inferred struct creation
	var myVar2 = myStruct{
		FirstName: "Mary",
	}

	fmt.Println(myVar2)
	fmt.Println(myVar2.getFirstName())

	// 3. Short struct creation
	myVar3 := myStruct{
		FirstName: "Jesus",
	}

	fmt.Println(myVar3)
	fmt.Println(myVar3.getFirstName())

	// 4. Creating a struct from imported function
	myVar4 := mypackage.NewMyStruct("Peter", "Gabriel")

	fmt.Println(myVar4)

	// 5. Changing a struct's field value using value receiver
	myVar3.changeCopyFirstName("Napoleon") // was Jesus
	fmt.Println(myVar3.getFirstName())     // still Jesus

	// 6. Changing a struct's field value using pointer receiver
	myVar3.changeFirstName("Napoleon") // was Jesus
	fmt.Println(myVar3.getFirstName()) // Napoleon
}
