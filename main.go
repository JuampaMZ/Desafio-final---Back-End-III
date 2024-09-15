package main

import (
	"log"
	"net/http"
	"odontology-appointments/db"
	"odontology-appointments/internal/appointment"
	"odontology-appointments/internal/dentist"
	"odontology-appointments/internal/patient"
	"odontology-appointments/internal/security"

	_ "odontology-appointments/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	db := db.InitDB()

	r := mux.NewRouter()

	// Dentist routes
	dentistRouter := r.PathPrefix("/dentists").Subrouter()
	dentistRouter.HandleFunc("/", dentist.GetAllDentists(db)).Methods("GET")
	dentistRouter.HandleFunc("/", security.Middleware(dentist.CreateDentist(db))).Methods("POST")
	dentistRouter.HandleFunc("/{id}", dentist.GetDentistByID(db)).Methods("GET")
	dentistRouter.HandleFunc("/{id}", security.Middleware(dentist.UpdateDentist(db))).Methods("PUT")
	dentistRouter.HandleFunc("/{id}", security.Middleware(dentist.PartialUpdateDentist(db))).Methods("PATCH")
	dentistRouter.HandleFunc("/{id}", security.Middleware(dentist.DeleteDentist(db))).Methods("DELETE")

	// Patient routes
	patientRouter := r.PathPrefix("/patients").Subrouter()
	patientRouter.HandleFunc("/", patient.GetAllPatients(db)).Methods("GET")
	patientRouter.HandleFunc("/", security.Middleware(patient.CreatePatient(db))).Methods("POST")
	patientRouter.HandleFunc("/{id}", patient.GetPatientByID(db)).Methods("GET")
	patientRouter.HandleFunc("/{id}", security.Middleware(patient.UpdatePatient(db))).Methods("PUT")
	patientRouter.HandleFunc("/{id}", security.Middleware(patient.PartialUpdatePatient(db))).Methods("PATCH")
	patientRouter.HandleFunc("/{id}", security.Middleware(patient.DeletePatient(db))).Methods("DELETE")

	// Appointment routes
	appointmentRouter := r.PathPrefix("/appointments").Subrouter()
	appointmentRouter.HandleFunc("/", appointment.GetAllAppointments(db)).Methods("GET")
	appointmentRouter.HandleFunc("/", security.Middleware(appointment.CreateAppointment(db))).Methods("POST")
	appointmentRouter.HandleFunc("/{id}", appointment.GetAppointmentByID(db)).Methods("GET")
	appointmentRouter.HandleFunc("/{id}", security.Middleware(appointment.UpdateAppointment(db))).Methods("PUT")
	appointmentRouter.HandleFunc("/{id}", security.Middleware(appointment.PartialUpdateAppointment(db))).Methods("PATCH")
	appointmentRouter.HandleFunc("/{id}", security.Middleware(appointment.DeleteAppointment(db))).Methods("DELETE")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
