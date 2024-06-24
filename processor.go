package main

import (
    "math"
    "regexp"
    "strconv"
    "strings"
    "time"
)

// Receipt represents a receipt with items
type Receipt struct {
    Retailer     string  `json:"retailer"`
    PurchaseDate string  `json:"purchaseDate"`
    PurchaseTime string  `json:"purchaseTime"`
    Items        []Item  `json:"items"`
    Total        string  `json:"total"`
}

// Item represents a line item in a receipt
type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price            string `json:"price"`
}

// calculatePoints processes the receipt and returns the total points
func calculatePoints(receipt Receipt) int {
    totalPoints := 0

    // One point for every alphanumeric character in the retailer name
    retailerChars := regexp.MustCompile(`[\w]`)
    totalPoints += len(retailerChars.FindAllString(receipt.Retailer, -1))

    // 50 points if the total is a round dollar amount with no cents
    if strings.HasSuffix(receipt.Total, ".00") {
        totalPoints += 50
    }

    // 25 points if the total is a multiple of 0.25
    totalFloat, _ := strconv.ParseFloat(receipt.Total, 64)
    if math.Mod(totalFloat, 0.25) == 0 {
        totalPoints += 25
    }

    // 5 points for every two items on the receipt
    totalPoints += (len(receipt.Items) / 2) * 5

    // Points for each item based on description length and price
    for _, item := range receipt.Items {
        if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
            itemPrice, _ := strconv.ParseFloat(item.Price, 64)
            pointsFromPrice := int(math.Ceil(itemPrice * 0.2))
            totalPoints += pointsFromPrice
        }
    }

    // 6 points if the day in the purchase date is odd
    date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
    if date.Day()%2 != 0 {
        totalPoints += 6
    }

    // 10 points if the time of purchase is between 2:00pm and 4:00pm
    time, _ := time.Parse("15:04", receipt.PurchaseTime)
    if time.Hour() >= 14 && time.Hour() < 16 {
        totalPoints += 10
    }

    return totalPoints
}