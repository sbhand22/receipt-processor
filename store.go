package main

import "sync"

// Store manages the in-memory storage of receipts
type Store struct {
    sync.Mutex
    Receipts map[string]int
}

// NewStore creates a new in-memory store
func NewStore() *Store {
    return &Store{
        Receipts: make(map[string]int),
    }
}

// AddReceipt saves a receipt's points with a unique ID
func (s *Store) AddReceipt(id string, points int) {
    s.Lock()
    defer s.Unlock()
    s.Receipts[id] = points
}

// GetPoints returns the points for a given receipt ID
func (s *Store) GetPoints(id string) (int, bool) {
    s.Lock()
    defer s.Unlock()
    points, exists := s.Receipts[id]
    return points, exists
}
