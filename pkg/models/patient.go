package models

type Patient struct {
	ID               int    `json:"id"`
	LastName         string `json:"last_name"`
	FirstName        string `json:"first_name"`
	Address          string `json:"address"`
	DNI              string `json:"dni"`
	RegistrationDate string `json:"registration_date"`
}
