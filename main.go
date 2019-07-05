package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello! It's %v", time.Now())
	})))
}
