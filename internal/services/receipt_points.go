package services

import (
	"log"
	"unicode"

	"receipt-processor/internal/models"
)

// 1. Alphanumeric chars in retailer name
func calcAlphaNumChars(s string) int {
	points := 0

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}

	return points
}

// 2. 50 pts if total is a round dollar amount (no cents)

// 3. 25 pts if total is a multiple of 0.25

// 4. 5 pts for every two items on receipt

// 5. If trimmed length of item description is a multiple length is a multiple of 3, multiple price by 0.2 and round up to nearest integer

// 6. 6 pts if day in the purchase date is odd

// 7. 10 pts if time of purchase is after 2 PM and before 4 PM

// Return point value of an entire receipt
func CalculatePoints(receipt models.Receipt) int {
	totalPoints := 0

	log.Printf("1. CalcAlphaNumChars = %v\n", calcAlphaNumChars(receipt.Retailer))
	totalPoints += calcAlphaNumChars(receipt.Retailer)

	log.Printf("Total Points = %v\n", totalPoints)
	return totalPoints
}
