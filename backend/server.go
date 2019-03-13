package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

func NewTx(writer http.ResponseWriter, request *http.Request) {
	if request.Body == nil {
		http.Error(writer, "No Body", 400)
		return
	}

	tx := &Tx{}
	err := json.NewDecoder(request.Body).Decode(&tx)

	if err != nil {
		http.Error(writer, err.Error(), 400)
		return
	}

	
	fmt.Println(tx)
}

func main() {
	http.Handle("/new-tx", corsHandler(NewTx))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func corsHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method != "OPTIONS" {
			h(w, r)
		}
	}
}
