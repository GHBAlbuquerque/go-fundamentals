package main

import (
	"log"
	"sort"
)

type Idol struct {
	StageName string
	GroupName string
}

func main() {
	// Maps
	myMap := make(map[string]string)

	myMap["dog"] = "Toby"
	myMap["other-dog"] = "Scooby"

	log.Println(myMap["dog"])

	myIdolMap := make(map[string]Idol)

	myIdolMap["Karina"] = Idol{StageName: "Karina", GroupName: "Aespa"}

	log.Println("This is", myIdolMap["Karina"].StageName, "from", myIdolMap["Karina"].GroupName, "!!")

	// Slices (similar to arrays)

	mySlice := []string{}

	mySlice = append(mySlice, "Winter", "NingNing", "Giselle")

	log.Println(mySlice)
	log.Println(mySlice[2])

	sort.Strings(mySlice)
	log.Println(mySlice)

	log.Println(mySlice[0:2])
}
