package main

import (
	"fmt"
	"github.com/VeRJiL/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Println("The server is running at localhost:8080/")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
