package main

// IWallet is the contract for interacting with a crypto
type IWallet interface {
	CreateAddress() string
	CheckBalance(string) float64
	Send(float64, string, string)
}
