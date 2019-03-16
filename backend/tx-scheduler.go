package main

import (
	"math/rand"
	"time"
)

// TxScheduler implements the ITxScheduler interface. Requires an ITxStatus
// to push payment updates too, and a IWallet to generate addresses and
// execute transactions
type TxScheduler struct {
	Statuses ITxStatus
	Wallet   IWallet
}

// ITxScheduler Given a TxSpec, schedule it, return a Tx to check the status
// Of that payment
type ITxScheduler interface {
	NewTxSpec(tx *TxSpec) *Tx
}

// NewTxSpec takes a TxSpec, creates a goroutine for the scheduling of the
// transactions, returns a Tx Object
func (ts *TxScheduler) NewTxSpec(spec *TxSpec) *Tx {
	id := randomString(10)
	wallet := ts.Wallet.CreateAddress()

	ts.Statuses.NewTx(id, len(spec.Outputs))
	tx := &Tx{
		DepositAddr: wallet,
		TxID:        id,
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
	ts.Statuses.Increment(tx.TxID)

	ts.Wallet.Send(balance, tx.DepositAddr, mixerAddr)

	sleepIntervals := nRandNumsThatSumToM(len(tx.Spec.Outputs), tx.Spec.Time)

	for index, output := range tx.Spec.Outputs {
		ts.Statuses.Increment(tx.TxID)
		time.Sleep(time.Duration(1000*sleepIntervals[index]) * time.Millisecond)
		amount := balance * (tx.Spec.Outputs[index].Split / 100)
		ts.Wallet.Send(amount, mixerAddr, output.Address)
	}

	ts.Statuses.Increment(tx.TxID)
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
