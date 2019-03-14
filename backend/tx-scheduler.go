package main

import (
	"fmt"
	"time"
)

type TxScheduler struct {
	Statuses ITxStatus
	Wallet   IWallet
}

type ITxScheduler interface {
	NewTxSpec(tx *TxSpec) *Tx
}

func (ts *TxScheduler) NewTxSpec(spec *TxSpec) *Tx {
	id := randomString(10)
	wallet := randomString(8)

	ts.Statuses.PushUpdate(id, "TX queued")
	tx := &Tx{
		DepositAddr: wallet,
		TxId:        id,
		Spec:        spec,
	}

	go ts.pollPayments(tx)
	return tx
}

func (ts *TxScheduler) pollPayments(tx *Tx) {
	for {
		fmt.Println(ts.Wallet.CheckBalance(tx.DepositAddr))
		time.Sleep(1000 * time.Millisecond)
	}
}
