package main

type TxSpec struct {
	Input   float64  `json:"input"`
	Outputs []Output `json:"outputs"`
	Time    float64  `json:"time"`
}

type Output struct {
	Address string  `json:"addr"`
	Split   float64 `json:"split"`
}
