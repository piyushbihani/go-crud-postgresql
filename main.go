package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/piyushbihani/go_stocks_crud/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on post 8080...")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
