package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/applied-concurrency-in-go/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to the Orders App!")
	handler, err := handlers.New()
	if err != nil {
		log.Fatal(err)
	}
	// start server
	router := ConfigureHandler(handler)
	fmt.Println("Listening on localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}

// ConfigureHandler configures the routes of this handler and binds handler functions to them
func ConfigureHandler(handler handlers.Handler) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/").
		Handler(http.HandlerFunc(handler.Index))
	router.Methods("GET").Path("/products").
		Handler(http.HandlerFunc(handler.ProductIndex))
	router.Methods("GET").Path("/orders/{orderId}").
		Handler(http.HandlerFunc(handler.OrderShow))
	router.Methods("POST").Path("/orders").
		Handler(http.HandlerFunc(handler.OrderInsert))

	return router
}
