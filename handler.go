package main

import (
    "encoding/json"
    "net/http"
    "github.com/google/uuid"
	"github.com/gorilla/mux"
)

var store = NewStore()

// ProcessReceipt handles the POST request to process a receipt
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Calculate points and store them
    points := calculatePoints(receipt)
    id := uuid.New().String()
    store.AddReceipt(id, points)

    json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// GetPoints handles the GET request to retrieve points for a given receipt ID
func GetPoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    points, exists := store.GetPoints(id)
    if !exists {
        http.Error(w, "No receipt found for that ID", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(map[string]int{"points": points})
}
