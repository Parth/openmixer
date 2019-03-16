package main

// Tx is the top level struct that encapsulates all the state related
// to a transaction
type Tx struct {
	TxID        string  `json:"id"`
	DepositAddr string  `json:"depositAddress"`
	Spec        *TxSpec `json:"spec"`
}
