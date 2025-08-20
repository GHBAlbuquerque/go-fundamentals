package mypackage

type myUnexportedStruct struct {
	FirstName  string
	secondName string
}

func NewMyStruct(name string, secondName string) myUnexportedStruct {
	return myUnexportedStruct{FirstName: name, secondName: secondName}
}
