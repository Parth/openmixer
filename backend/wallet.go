package main

type IWallet interface {
	CreateAddress() string
	CheckBalance(string) float64
	Send(float64, string, string)
}