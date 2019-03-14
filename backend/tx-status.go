package main

import "sync"

type TxStatus struct {
	*sync.RWMutex

	Statuses map[string]string
}

type ITxStatus interface {
	PushUpdate(key string, update string)
}

func (ts *TxStatus) PushUpdate(key string, update string) {
	ts.Lock()
	defer ts.Unlock()

	current := ts.Statuses[key]
	newUpdate := current + "\n" + update
	ts.Statuses[key] = newUpdate
}
