package main

import (
	"enterprise-computing/tracks/repository"
	"enterprise-computing/tracks/resources"
	"log"
	"net/http"
)

func main() {
	repository.Init()
	repository.Create()
	log.Fatal(http.ListenAndServe(":3000", resources.Router()))
}
