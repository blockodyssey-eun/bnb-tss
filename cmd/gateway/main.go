package main

import (
	"log"
	"tss_project/internal/gateway"
)

func main() {
	server := gateway.NewServer()
	log.Println("Starting Gateway server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
