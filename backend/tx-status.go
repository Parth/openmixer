package main

import (
	"fmt"
	"sync"
)

type TxStatus struct {
	sync.RWMutex

	Statuses map[string][]string
}

type ITxStatus interface {
	PushUpdate(key string, update string)
	GetUpdates(key string) []string
}

func (ts *TxStatus) PushUpdate(key string, update string) {
	ts.Lock()
	defer ts.Unlock()
	fmt.Println(key, update)

	ts.Statuses[key] = append(ts.Statuses[key], update)
}

func (ts *TxStatus) GetUpdates(key string) []string {
	ts.Lock()
	defer ts.Unlock()

	return ts.Statuses[key]
}
