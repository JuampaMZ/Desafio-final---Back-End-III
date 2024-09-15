package models

type Appointment struct {
	ID          int    `json:"id"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Description string `json:"description"`
	PatientID   int    `json:"patient_id"`
	DentistID   int    `json:"dentist_id"`
}
