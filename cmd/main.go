package main

import (
	. "../handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	log.Printf("Server is started ")
	defer log.Printf("Server is stopped ")
	r := mux.NewRouter()
	r.HandleFunc("/api/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products", PostProductHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}", PutProductHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id}", DeleteProductHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
}
