package main

type TxSpec struct {
	Input   float64  `json:"input"`
	Outputs []string `json: "outputs"`
	Time    float64  `json: "time"`
}
