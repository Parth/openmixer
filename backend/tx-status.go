package main

import (
	"errors"
	"sync"
)

type Status struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type TxStatus struct {
	sync.RWMutex

	Statuses map[string]*Status
}

type ITxStatus interface {
	NewTx(key string, total int)
	Increment(key string) error
	GetStatus(key string) *Status
}

func (ts *TxStatus) NewTx(key string, total int) {
	ts.Lock()
	defer ts.Unlock()

	s := &Status{
		Current: -1,
		Total:   total,
	}
	ts.Statuses[key] = s
}

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

func (ts *TxStatus) GetStatus(key string) *Status {
	ts.Lock()
	defer ts.Unlock()

	return ts.Statuses[key]
}
