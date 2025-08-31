package models

type Admin struct {
	Person      Person
	Username    string `json:"Username"`
	Permissions string `json:"Permissions"`
}
