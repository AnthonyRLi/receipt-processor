# receipt-processor

Language : Go

To Use:

1. Assuming you have Go installed, once you have git pulled, run: go mod tidy

2. From receipt-processor folder, run: go run cmd/server/main.go

3. Will open on Port 3000

4. Using Postman:

    POST Route: 
    
        http://localhost:3000/receipts/process

    GET Route: 
    
        http://localhost:3000/receipts/{id}/points


5. Using curl on a terminal:

    POST:

        curl -X POST \
        -H "Content-Type: application/json" \
        -d '{
            "retailer": "Target",
            "purchaseDate": "2022-01-01",
            "purchaseTime": "13:01",
            "items": [
            {
                "shortDescription": "Mountain Dew 12PK",
                "price": "6.49"
            },
            {
                "shortDescription": "Emils Cheese Pizza",
                "price": "12.25"
            },
            {
                "shortDescription": "Knorr Creamy Chicken",
                "price": "1.26"
            },
            {
                "shortDescription": "Doritos Nacho Cheese",
                "price": "3.35"
            },
            {
                "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
                "price": "12.00"
            }
            ],
            "total": "35.35"
        }' \
        http://localhost:3000/receipts/process


    GET:

        curl http://localhost:3000/receipts/{id}/points




Thanks!
Anthony