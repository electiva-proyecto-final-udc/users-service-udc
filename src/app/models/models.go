package models

import "time"

type Person struct {
	ID           string `json:"ID"`
	DocumentType string `json:"DocumentType"`
	Document     string `json:"Document"`
	Name         string `json:"Name"`
	Surname      string `json:"Surname"`
	Email        string `json:"Email"`
	PhoneNumber  string `json:"PhoneNumber"`
}

type Admin struct {
	Person      Person
	Username    string `json:"Username"`
	Permissions string `json:"Permissions"`
}

type Client struct {
	Person  Person
	Address string `json:"Address"`
}

type Tecnician struct {
	Person    Person
	Username  string    `json:"Username"`
	Address   string    `json:"Address"`
	IsActive  bool      `json:"IsActive"`
	EntryDate time.Time `json:"EntryDate"`
}
