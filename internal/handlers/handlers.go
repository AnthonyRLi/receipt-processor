package handlers

import (
	"encoding/json"
	"net/http"

	"receipt-processor/internal/models"
	"receipt-processor/internal/services"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// In Memory Storage
var receiptPoints = make(map[string]int64)

// POST
func PostReceiptHandler(writer http.ResponseWriter, request *http.Request) {

	// Decode to JSON
	var receipt models.Receipt
	if err := json.NewDecoder(request.Body).Decode(&receipt); err != nil {
		http.Error(writer, "Invalid Request Body.", http.StatusBadRequest)
		return
	}

	// Generate unique UUID for receipt
	id := uuid.New().String()

	// Calculate receipt points
	points := services.CalculatePoints(receipt)

	// Store receipt points
	// Could do error handling in case there is somehow a UUID overlap
	receiptPoints[id] = points

	// Return response
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(models.PostResponse{ID: id})

}

// GET
func GetReceiptHandler(writer http.ResponseWriter, request *http.Request) {

	// Get UUID from params
	params := mux.Vars(request)
	id := params["id"]

	// Check if receipt exists
	points, valid := receiptPoints[id]
	if !valid {
		http.Error(writer, "Invalid receipt ID.", http.StatusBadRequest)
		return
	}

	// Return response
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(models.GetResponse{Points: points})

}
