package main

import (
	"go-packages/helpers"
	"log"
)

func main() {
	exportedStruct := helpers.MyExportedStruct{
		Val: "anything",
	}

	log.Println(exportedStruct)
	log.Println(helpers.MyExportedFunc())
}
