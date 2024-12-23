package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

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
		http.Error(writer, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	// Run validator on Receipt to see if it is valid.
	errMsg := models.ReceiptValid(receipt)
	if errMsg != nil {
		log.Printf("Error: %v", errMsg.Message)
		http.Error(writer, "The receipt is invalid.", http.StatusBadRequest)
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

	// Check Regex for UUID
	// Keeping same error message for now as I don't see a requirement for what you would return in the case of a request for an invalid receipt ID. Naturally, you would get the NotFound, so maybe not needed?
	// Only consideration is that you would want to reject potential attackers trying to go for SQL injection

	if !regexp.MustCompile(`^\S+$`).MatchString(id) {
		http.Error(writer, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	// Check if receipt exists
	points, valid := receiptPoints[id]
	if !valid {
		http.Error(writer, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	// Return response
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(models.GetResponse{Points: points})

}
