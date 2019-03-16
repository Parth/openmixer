package main

import (
	"testing"
)

func TestTxStatus(t *testing.T) {
	s := &TxStatus{
		Statuses: map[string]*Status{},
	}

	if s.GetStatus("test") != nil {
		t.Errorf("key not found should be nil should be nil")
	}

	err := s.Increment("test")
	if err == nil {
		t.Errorf("should throw error if trying to incrememnt not-known-key")
	}

	s.NewTx("test", 3)

	if s.GetStatus("test").Current != -1 {
		t.Errorf("Should start with -1")
	}

	s.Increment("test")
	if s.GetStatus("test").Current != 0 {
		t.Errorf("Increment should work")
	}

	if s.GetStatus("test").Total != 3 {
		t.Errorf("was passed 3 as total")
	}
}
