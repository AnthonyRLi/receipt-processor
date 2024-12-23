package handlers

import (
	"encoding/json"
	"net/http"

	"receipt-processor/internal/models"

	"github.com/google/uuid"
)

// In Memory Storage
var receiptPoints = make(map[string]int)

// POST
func PostReceipthandler(writer http.ResponseWriter, request *http.Request) {

	// Decode to JSON
	var receipt models.Receipt
	if err := json.NewDecoder(request.Body).Decode(&receipt); err != nil {
		http.Error(writer, "Invalid Request Body.", http.StatusBadRequest)
		return
	}

	// Generate unique UUID for receipt
	id := uuid.New().String()

	// Calculate receipt points
	var points int = 0

	// Store receipt points
	// Could do error handling in case there is somehow a UUID overlap
	receiptPoints[id] = points

	// Return response
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(models.PostResponse{ID: id})

}

// GET
