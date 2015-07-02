package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening.")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3002", nil)

}
