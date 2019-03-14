package main

type Tx struct {
	TxId        string  `json:"id"`
	DepositAddr string  `json: "deposit-addr"`
	Spec        *TxSpec `json: "spec"`
}
