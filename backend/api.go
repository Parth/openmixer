package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// API is what infrastructure layer components will call to interact with our business logic
type API struct {
	Scheduler ITxScheduler
	Statuses  ITxStatus
	Wallet    IWallet
}

// NewTx parses http requests, validates them, and gives the transaction described
// to the transaction scheduler. The transaction scheduler returns a txid which a
// client can use to keep track of the status of the transaction
func (a *API) NewTx(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, "No Body", 400)
		return
	}

	spec := &TxSpec{}
	err := json.NewDecoder(request.Body).Decode(&spec)
	fmt.Println(spec)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	tx := a.Scheduler.NewTxSpec(spec)

	fmt.Println(tx)
	json.NewEncoder(writer).Encode(tx)
}

// TxStatus parses http requests, validates them, and looks up the status of the
// transaction in question returns a Status object
func (a *API) TxStatus(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, "No Body", 400)
		return
	}

	txidMap := map[string]string{}
	err := json.NewDecoder(request.Body).Decode(&txidMap)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	json.NewEncoder(writer).Encode(a.Statuses.GetStatus(txidMap["id"]))
}
