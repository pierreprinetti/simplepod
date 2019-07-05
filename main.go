package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Hello, world!"))
	})))
}
