package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()

	log.Printf(fmt.Sprintf("Service listening on port %v", os.Getenv("PORT")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), router))
}
