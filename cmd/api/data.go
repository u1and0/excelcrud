package api

import (
	"time"
)

type (
	// Datum : Excel シート上の一行
	Datum struct {
		EntryDate  time.Time `json:"entrydate"`
		UserID     string    `json:"userid"`
		Name       string    `json:"name"`
		Sex        string    `json:"sex"`
		Age        int       `json:"age"`
		TotalMoney int       `json:"totalmoney"`
		BirthDay   time.Time `json:"birthday"`
	}
	// Data : Excel シート上の全行
	Data []Datum
)

// New : Datum constructor
func New() *Datum {
	return &Datum{}
}
