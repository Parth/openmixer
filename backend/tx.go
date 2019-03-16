package main

type Tx struct {
	TxId        string  `json:"id"`
	DepositAddr string  `json:"depositAddress"`
	Spec        *TxSpec `json:"spec"`
}
