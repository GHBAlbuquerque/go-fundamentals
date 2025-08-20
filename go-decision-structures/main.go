package main

import "log"

func main() {
	var isFalse bool
	isFalse = false
	isTrue := true
	cat := "cat"

	if cat == "cat" && isTrue {
		log.Println(isTrue)
	} else if cat == "cat" && isFalse {
		log.Println(isFalse)
	} else {
		log.Println(isFalse)
	}

	animal := "horse"

	switch animal {
	case "cat":
		log.Println("animal is a cat")
	case "dog":
		log.Println("animal is a dog")
	case "horse":
		log.Println("animal is a horse")
	}

}
