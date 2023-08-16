package main

import (
	"fmt"
	"net/http"

	"github.com/faizshaikh22/bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	serverAddr := "localhost:9010"

	fmt.Printf("Server is starting on http://%s\n", serverAddr)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
