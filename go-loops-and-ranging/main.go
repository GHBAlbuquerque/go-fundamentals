package main

import "log"

func main() {

	// classic for
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "cat", "horse", "cow"}

	// for with resources
	for i, animal := range animals {
		log.Println(i, animal)
	}

	// for with resources, ignores index
	for _, animal := range animals {
		log.Println(animal)
	}

	idols := map[string]string{}
	idols["karina"] = "aespa"
	idols["winter"] = "aespa"
	idols["ningning"] = "aespa"
	idols["giselle"] = "aespa"

	// for with map -> range returns key and value
	for key, value := range idols {
		log.Println(key)
		log.Println(value)
	}

	var firstLine = "once upon a time"

	// for with string -> range returns index and a rune (byte32 that represents a char)
	for i, l := range firstLine {
		log.Printf("Rune %d as a character: %c\n", i, l)
	}
}
