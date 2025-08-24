package main

import (
	"fmt"
	"net/http"
)

func main() {
	//built in package of the standard library
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(writer, "Hello, world!")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Bytes written:", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
