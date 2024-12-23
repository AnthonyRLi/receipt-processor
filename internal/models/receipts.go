package models

import (
	"regexp"
	"time"
)

// Receipt
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item on receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// POST response with Receipt ID
type PostResponse struct {
	ID string `json:"id"`
}

// GET response with Receipt ID
type GetResponse struct {
	Points int64 `json:"points"`
}

// Error delivery struct
type ErrorMessage struct {
	IsError bool
	Message string
}

// Create and return error
func ReturnError(isError bool, message string) *ErrorMessage {
	return &ErrorMessage{
		IsError: isError,
		Message: message,
	}
}

// Item Validator
func ItemValid(item Item) *ErrorMessage {
	// Check Item Description
	if !regexp.MustCompile(`^[\w\s\-]+$`).MatchString(item.ShortDescription) {
		return ReturnError(true, "Retailer Name has an invalid character.")
	}

	// Check Total regex
	if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(item.Price) {
		return ReturnError(true, "Total has an invalid character.")
	}

	return nil
}

// Receipt Validator
func ReceiptValid(receipt Receipt) *ErrorMessage {

	// Check if all fields are included
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || receipt.Total == "" || len(receipt.Items) < 1 {
		return ReturnError(true, "Missing a required field.")
	}

	// Check Retailer regex
	if !regexp.MustCompile(`^[\w\s\-&]+$`).MatchString(receipt.Retailer) {
		return ReturnError(true, "Retailer Name has an invalid character.")
	}

	// Check Total regex
	if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(receipt.Total) {
		return ReturnError(true, "Total has an invalid character.")
	}

	// Check Purchase Date
	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		return ReturnError(true, "Purchase Date is in invalid format.")
	}

	// Check Purchase Time
	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return ReturnError(true, "Purchase Time is in invalid format.")
	}

	// Check All Items
	for _, item := range receipt.Items {
		if errMsg := ItemValid(item); errMsg != nil {
			return errMsg
		}
	}

	return nil
}
