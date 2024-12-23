package services

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"receipt-processor/internal/models"
)

// 1. Alphanumeric chars in retailer name
// One point for every alphanumeric character in the retailer name.
func calcAlphaNumChars(retailerName string) int64 {
	points := 0

	for _, char := range retailerName {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}

	log.Printf("1. Calc AlphaNum Chars = %v\n", points)
	return int64(points)
}

// 2. 50 pts if total is a round dollar amount (no cents)
// 50 points if the total is a round dollar amount with no cents.
func roundDollarTotal(totalPrice string) int64 {
	points := 0
	total, _ := strconv.ParseFloat(totalPrice, 64)

	// Check if there is a decimal remainder by moduloing by 1
	if math.Mod(total, 1.0) == 0.0 {
		points += 50
	}

	log.Printf("2. Round Dollar Total = %v\n", points)
	return int64(points)
}

// 3. 25 pts if total is a multiple of 0.25
// 25 points if the total is a multiple of 0.25.
func multipleOfQuarter(totalPrice string) int64 {
	points := 0
	total, _ := strconv.ParseFloat(totalPrice, 64)

	// Check if there is a decimal remainder by moduloing by 0.25
	if math.Mod(total, 0.25) == 0.0 {
		points += 25
	}

	log.Printf("3. Total Multiple Of 0.25 = %v\n", points)
	return int64(points)
}

// 4. 5 pts for every two items on receipt
// 5 points for every two items on the receipt.
func everyTwoItems(items []models.Item) int64 {
	points := 0
	points = len(items) / 2 * 5

	log.Printf("4. Every Two Items = %v\n", points)
	return int64(points)
}

// 5. If trimmed length of item description is a multiple length is a multiple of 3, multiply price by 0.2 and round up to nearest integer
// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
func itemDescriptionLength(items []models.Item) int64 {
	points := 0

	for _, item := range items {
		// Trim spaces out of item description
		trimmedLen := len(strings.TrimSpace(item.ShortDescription))

		// Check if divisible by 3
		// If so, multiply price by 0.2 and round up, add to points
		if trimmedLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	log.Printf("5. Item Description Length = %v\n", points)
	return int64(points)
}

// 6. 6 pts if day in the purchase date is odd
// 6 points if the day in the purchase date is odd.
func oddDate(purchaseDate string) int64 {
	points := 0
	parsedPurchaseDate, _ := time.Parse("2006-01-02", purchaseDate)

	// 0 points if date is an odd number
	if parsedPurchaseDate.Day()%2 != 0 {
		points += 6
	}

	log.Printf("6. Odd Date = %v\n", points)
	return int64(points)
}

// 7. 10 pts if time of purchase is after 2 PM and before 4 PM
// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func betweenTime(purchaseTime string) int64 {
	points := 0

	parsedPurchaseTime, _ := time.Parse("15:04", purchaseTime)
	beginning, _ := time.Parse("15:04", "14:00")
	end, _ := time.Parse("15:04", "16:00")

	// Return 10 points if after 2 PM and before 4 PM
	if parsedPurchaseTime.After(beginning) && parsedPurchaseTime.Before(end) {
		points += 10
	}

	log.Printf("7. Between 2 and 4PM  = %v\n", points)
	return int64(points)
}

// Return point value of an entire receipt
func CalculatePoints(receipt models.Receipt) int64 {
	log.Printf("\nReceipt: %v", receipt.Retailer)

	totalPoints := int64(0)

	totalPoints += calcAlphaNumChars(receipt.Retailer)
	totalPoints += roundDollarTotal(receipt.Total)
	totalPoints += multipleOfQuarter(receipt.Total)
	totalPoints += everyTwoItems(receipt.Items)
	totalPoints += itemDescriptionLength(receipt.Items)
	totalPoints += oddDate(receipt.PurchaseDate)
	totalPoints += betweenTime(receipt.PurchaseTime)

	log.Printf("\nTotal Points = %v\n", totalPoints)
	return totalPoints
}
