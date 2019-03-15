package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Api struct {
	Scheduler ITxScheduler
	Statuses  ITxStatus
	Wallet    IWallet
}

// NewTx parses http requests, validates them, and gives the transaction described
// to the transaction scheduler. The transaction scheduler returns a txid which a
// client can use to keep track of the status of the transaction
func (a *Api) NewTx(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, "No Body", 400)
		return
	}

	spec := &TxSpec{}
	err := json.NewDecoder(request.Body).Decode(&spec)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	tx := a.Scheduler.NewTxSpec(spec)
	fmt.Println(tx)
	json.NewEncoder(writer).Encode(tx)
}

func (a *Api) TxStatus(writer http.ResponseWriter, request *http.Request) {
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

	json.NewEncoder(writer).Encode(a.Statuses.GetUpdates(txidMap["id"]))
}
