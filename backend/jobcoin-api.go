package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type JobcoinAPI struct {
	URL string
}

func (j *JobcoinAPI) CreateAddress() string {
	return ""
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

func (j *JobcoinAPI) Send(float64, string, string) {}
