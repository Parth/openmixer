package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// JobcoinAPI is an implementation of the Wallet interface
type JobcoinAPI struct {
	URL string
}

// CreateAddress generates a new Jobcoin wallet, used as a unique
// deposit location for every customer
func (j *JobcoinAPI) CreateAddress() string {
	return randomString(8)
}

// CheckBalance returns the balance for a given Jobcoin address
func (j *JobcoinAPI) CheckBalance(addr string) float64 {
	resp, err := http.Get(j.URL + "addresses/" + addr)
	if err != nil {
		fmt.Println("TODO") // TODO
	}

	responseMap := map[string]string{}

	json.NewDecoder(resp.Body).Decode(&responseMap)
	balance, err := strconv.ParseFloat(responseMap["balance"], 64)
	if err != nil {
		return -1
	}

	return balance
}

// Send creates and executes a transaction
func (j *JobcoinAPI) Send(amount float64, from string, to string) {
	message := map[string]string{
		"fromAddress": from,
		"toAddress":   to,
		"amount":      strconv.FormatFloat(amount, 'f', 6, 64),
	}

	fmt.Println(message)

	jsonValue, err := json.Marshal(message)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(http.Post(
		j.URL+"transactions",
		"application/json",
		bytes.NewBuffer(jsonValue)))
}
