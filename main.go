package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/receipts/process", ProcessReceipt).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", GetPoints).Methods("GET")

    // Set up the server
    http.Handle("/", r)
    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
