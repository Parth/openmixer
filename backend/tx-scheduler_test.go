package main

import (
	"strconv"
	"testing"
	"time"
)

type MockWallet struct {
	addressIndex int
	timesChecked int
	transactions map[string]float64
}

func (mw *MockWallet) CreateAddress() string {
	mw.addressIndex++
	return "address" + strconv.Itoa(mw.addressIndex)
}

func (mw *MockWallet) CheckBalance(addr string) float64 {
	if mw.timesChecked == 0 {
		mw.timesChecked++
		return 0
	} else if mw.timesChecked == 1 {
		mw.timesChecked++
		return 1
	} else {
		mw.timesChecked++
		return 10
	}
}

type MockStatus struct {
}

func (ms *MockStatus) NewTx(string, int)        {}
func (ms *MockStatus) Increment(string) error   { return nil }
func (ms *MockStatus) GetStatus(string) *Status { return nil }

func (mw *MockWallet) Send(amount float64, from string, to string) {
	mw.transactions[from] += (-1 * amount)
	mw.transactions[to] += amount
}

func TestNewTxSpec(t *testing.T) {
	mockWallet := &MockWallet{
		transactions: map[string]float64{},
	}

	mockStatus := &MockStatus{}

	scheduler := &TxScheduler{
		Statuses: mockStatus,
		Wallet:   mockWallet,
	}

	scheduler.NewTxSpec(&TxSpec{
		Input: 8,
		Time:  3,
		Outputs: []Output{
			Output{
				Address: "test1",
				Split:   30,
			},
			Output{
				Address: "test2",
				Split:   30,
			},
			Output{
				Address: "test3",
				Split:   40,
			},
		},
	})

	time.Sleep(5 * time.Second) // This makes unit tests sad

	if mockWallet.transactions["address1"] != -10 {
		t.Errorf("address1 should have -10")
	}

	if mockWallet.transactions["test1"] != 0.3*10 {
		t.Errorf("expected %f, got %f", 10*0.3, mockWallet.transactions["test1"])
	}

	if mockWallet.transactions["test3"] != 0.4*10 {
		t.Errorf("expected %f, got %f", 10*0.4, mockWallet.transactions["test1"])
	}

	if mockWallet.transactions[mixerAddr] != 0 {
		t.Errorf("house mixer should be empty")
	}
}
