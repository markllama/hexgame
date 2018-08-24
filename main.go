package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Goodbye")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8000", nil)
}
