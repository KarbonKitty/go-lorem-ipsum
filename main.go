package main

import (
	"fmt"
	"net/http"
)

func loremIpsum(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "lorem ipsum dolor sit amet")
}

func main() {
	http.HandleFunc("/lorem-ipsum", loremIpsum)

	http.ListenAndServe(":8080", nil)
}
