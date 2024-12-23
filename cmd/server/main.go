package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"receipt-processor/internal/handlers"
)

func main() {
	router := mux.NewRouter()

	const port string = "3000"

	router.HandleFunc("/receipts/process", handlers.PostReceipthandler).Methods("POST")

	log.Printf("Server started on Port %s:", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}

}
