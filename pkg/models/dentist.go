package models

type Dentist struct {
	ID        int    `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	License   string `json:"license"`
}
