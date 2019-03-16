package main

// TxSpec describes a new transaction on this platform
type TxSpec struct {
	// Deposit amount that triggers the transactions
	Input float64 `json:"input"`

	// Where will the money & how much
	Outputs []Output `json:"outputs"`

	// By when should the last payment be received
	Time float64 `json:"time"`
}

// Output describes where the money will be sent
// and what percentage of the total transaction this
// output is
type Output struct {
	// Destination
	Address string `json:"addr"`

	// [0-100]% of the balance post Input detection
	Split float64 `json:"split"`
}
