package models

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
	Points int `json:"points"`
}
