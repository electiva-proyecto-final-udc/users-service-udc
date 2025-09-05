package models

type Admin struct {
	ID           string `json:"ID"`
	DocumentType string `json:"DocumentType"`
	Document     string `json:"Document"`
	Name         string `json:"Name"`
	Surname      string `json:"Surname"`
	Email        string `json:"Email"`
	PhoneNumber  string `json:"PhoneNumber"`
	Username     string `json:"Username"`
	Password     string `json:"Password"`
	Permissions  string `json:"Permissions"`
}
