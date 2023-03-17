package main

import (
	"enterprise-computing/search/resources"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":3001", resources.Router()))
}
