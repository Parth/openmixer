package main

import (
	"math/rand"
	"strconv"
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
	wallet := ts.Wallet.CreateAddress()

	ts.Statuses.PushUpdate(id, "TX queued")
	tx := &Tx{
		DepositAddr: wallet,
		TxId:        id,
		Spec:        spec,
	}

	go ts.newPaymentsWorker(tx)
	return tx
}

func (ts *TxScheduler) newPaymentsWorker(tx *Tx) {

	for currentBalance := ts.Wallet.CheckBalance(tx.DepositAddr); currentBalance < tx.Spec.Input; currentBalance = ts.Wallet.CheckBalance(tx.DepositAddr) {
		time.Sleep(500 * time.Millisecond)
	}

	balance := ts.Wallet.CheckBalance(tx.DepositAddr)
	ts.Statuses.PushUpdate(tx.TxId, "Deposit detected")

	ts.Wallet.Send(balance, tx.DepositAddr, mixerAddr)

	sleepIntervals := nRandNumsThatSumToM(len(tx.Spec.Outputs), tx.Spec.Time)

	for index, outputAddr := range tx.Spec.Outputs {
		ts.Statuses.PushUpdate(tx.TxId, "Payment #"+strconv.Itoa(index)+" scheduled")
		time.Sleep(time.Duration(1000*sleepIntervals[index]) * time.Millisecond)
		amount := balance * (tx.Spec.Splits[index] / 100)
		ts.Wallet.Send(amount, mixerAddr, outputAddr)
	}
}

// TODO move to random utils
func nRandNumsThatSumToM(n int, m float64) []float64 {
	nums := make([]float64, n)
	sum := 0.0

	for i := 0; i < n; i++ {
		num := rand.Float64()
		sum += num
		nums[i] = num
	}

	factor := m / sum

	for index := range nums {
		nums[index] *= factor
	}

	return nums
}
