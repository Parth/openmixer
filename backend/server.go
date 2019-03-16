package main

import (
	"log"
	"net/http"
)

const mixerAddr = "transparent-mixer"

func main() {

	wallet := &JobcoinAPI{
		URL: "http://jobcoin.gemini.com/sanitary/api/",
	}

	statuses := &TxStatus{
		Statuses: map[string]*Status{},
	}

	scheduler := &TxScheduler{
		Wallet:   wallet,
		Statuses: statuses,
	}

	api := &Api{
		Scheduler: scheduler,
		Statuses:  statuses,
		Wallet:    wallet,
	}

	http.Handle("/new-tx", corsHandler(api.NewTx))
	http.Handle("/tx-status", corsHandler(api.TxStatus))

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
