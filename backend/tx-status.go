package main

import (
	"errors"
	"sync"
)

// Status where has the scheduler reached?
// Current: -1 means no deposit, [0-total] which transaction is being scheduled, total+1 indicates we're done
type Status struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

// TxStatus a light wrapper around map[string]*Status for threadsafety,
// Each http request spawns it's own goroutine
type TxStatus struct {
	sync.RWMutex

	Statuses map[string]*Status
}

// ITxStatus contract for getting & setting updates for a transaction
type ITxStatus interface {
	NewTx(key string, total int)
	Increment(key string) error
	GetStatus(key string) *Status
}

// NewTx associates the given key with a starting status of -1 (no deposit)
func (ts *TxStatus) NewTx(key string, total int) {
	ts.Lock()
	defer ts.Unlock()

	s := &Status{
		Current: -1,
		Total:   total,
	}
	ts.Statuses[key] = s
}

// Increment increases the status.Current of the associated key, or
// returns an error if that key does not exist
func (ts *TxStatus) Increment(key string) error {
	ts.Lock()
	defer ts.Unlock()

	_, exists := ts.Statuses[key]
	if !exists {
		return errors.New("key does not exist")
	}
	ts.Statuses[key].Current++
	return nil
}

// GetStatus returns the associated status
func (ts *TxStatus) GetStatus(key string) *Status {
	ts.Lock()
	defer ts.Unlock()

	return ts.Statuses[key]
}
