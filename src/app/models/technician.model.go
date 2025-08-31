package models

import "time"

type Technician struct {
	Person    Person
	Username  string    `json:"Username"`
	Address   string    `json:"Address"`
	IsActive  bool      `json:"IsActive"`
	EntryDate time.Time `json:"EntryDate"`
}
