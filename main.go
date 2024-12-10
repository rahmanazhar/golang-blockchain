package main

import (
	"blockchain/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := api.NewBlockchainServer()
	handler := server.SetupRoutes()

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
