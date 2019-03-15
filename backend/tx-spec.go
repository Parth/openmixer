package main

type TxSpec struct {
	Input   float64   `json:"input"`
	Outputs []string  `json: "outputs"`
	Splits  []float64 `json: "splits"`
	Time    float64   `json: "time"`
}
