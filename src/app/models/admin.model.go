package models

type Admin struct {
	ID           string `json:"id"`
	DocumentType string `json:"documentType"`
	Document     string `json:"document"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Permissions  string `json:"permissions"`
}
