package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type JobcoinAPI struct {
	URL string
}

func (j *JobcoinAPI) CreateAddress() string {
	return randomString(8)
}

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

func (j *JobcoinAPI) Send(amount float64, from string, to string) {
	message := map[string]string{
		"fromAddress": from,
		"toAddress":   to,
		"amount":      strconv.FormatFloat(amount, 'f', 6, 64),
	}

	jsonValue, _ := json.Marshal(message)

	http.Post(
		j.URL+"transactions",
		"application/json",
		bytes.NewBuffer(jsonValue))
}
