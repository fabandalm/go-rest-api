package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprint(w, "HelloWorld!")
	})

	http.ListenAndServe("localhost:8000", nil)

}

