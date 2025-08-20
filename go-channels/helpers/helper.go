package helpers

import "math/rand"

func RandomNumber(n int) int {
	value := rand.Intn(n) // function that generates a random number within the size of n
	return value
}
