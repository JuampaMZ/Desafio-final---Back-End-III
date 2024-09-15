package patient

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"odontology-appointments/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

// GET: Obtener todos los pacientes
// @Summary Listar todos los pacientes
// @Tags Paciente
// @Produce json
// @Success 200 {array} models.Patient
// @Router /pacientes [get]
func GetAllPatients(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM patients")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var patients []models.Patient
		for rows.Next() {
			var patient models.Patient
			rows.Scan(&patient.ID, &patient.LastName, &patient.FirstName, &patient.Address, &patient.DNI, &patient.RegistrationDate)
			patients = append(patients, patient)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patients)
	}
}

// POST: Crear un nuevo paciente
// @Summary Agregar un nuevo paciente
// @Tags Paciente
// @Accept json
// @Produce json
// @Param paciente body models.Patient true "Paciente"
// @Success 201 {object} models.Patient
// @Router /pacientes [post]
func CreatePatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var patient models.Patient
		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("INSERT INTO patients (last_name, first_name, address, dni, registration_date) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res, err := stmt.Exec(patient.LastName, patient.FirstName, patient.Address, patient.DNI, patient.RegistrationDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, _ := res.LastInsertId()
		patient.ID = int(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patient)
	}
}

// GET: Obtener paciente por ID
// @Summary Obtener un paciente por ID
// @Tags Paciente
// @Produce json
// @Param id path int true "ID del paciente"
// @Success 200 {object} models.Patient
// @Failure 404 {object} models.Error
// @Router /pacientes/{id} [get]
func GetPatientByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var patient models.Patient
		err = db.QueryRow("SELECT * FROM patients WHERE id = ?", id).Scan(
			&patient.ID, &patient.LastName, &patient.FirstName, &patient.Address, &patient.DNI, &patient.RegistrationDate)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Patient not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patient)
	}
}

// PUT: Actualizar paciente
// @Summary Actualizar un paciente
// @Tags Paciente
// @Accept json
// @Produce json
// @Param id path int true "ID del paciente"
// @Param paciente body models.Patient true "Paciente"
// @Success 200 {object} models.Patient
// @Failure 404 {object} models.Error
// @Router /pacientes/{id} [put]
func UpdatePatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var patient models.Patient
		err = json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("UPDATE patients SET last_name = ?, first_name = ?, address = ?, dni = ?, registration_date = ? WHERE id = ?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = stmt.Exec(patient.LastName, patient.FirstName, patient.Address, patient.DNI, patient.RegistrationDate, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(patient)
	}
}

// PATCH: Actualizar parcialmente paciente
// @Summary Actualizar algunos campos de un paciente
// @Tags Paciente
// @Accept json
// @Produce json
// @Param id path int true "ID del paciente"
// @Success 200 {object} models.Patient
// @Failure 404 {object} models.Error
// @Router /pacientes/{id} [patch]
func PartialUpdatePatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var fields map[string]interface{}
		err = json.NewDecoder(r.Body).Decode(&fields)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := "UPDATE patients SET "
		args := []interface{}{}
		first := true

		if lastName, ok := fields["last_name"].(string); ok {
			if !first {
				query += ", "
			}
			query += "last_name = ?"
			args = append(args, lastName)
			first = false
		}

		if firstName, ok := fields["first_name"].(string); ok {
			if !first {
				query += ", "
			}
			query += "first_name = ?"
			args = append(args, firstName)
			first = false
		}

		if address, ok := fields["address"].(string); ok {
			if !first {
				query += ", "
			}
			query += "address = ?"
			args = append(args, address)
			first = false
		}

		if dni, ok := fields["dni"].(string); ok {
			if !first {
				query += ", "
			}
			query += "dni = ?"
			args = append(args, dni)
			first = false
		}

		if registrationDate, ok := fields["registration_date"].(string); ok {
			if !first {
				query += ", "
			}
			query += "registration_date = ?"
			args = append(args, registrationDate)
		}

		query += " WHERE id = ?"
		args = append(args, id)

		stmt, err := db.Prepare(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = stmt.Exec(args...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// DELETE: Eliminar paciente
// @Summary Eliminar un paciente
// @Tags Paciente
// @Produce json
// @Param id path int true "ID del paciente"
// @Success 204
// @Failure 404 {object} models.Error
// @Router /pacientes/{id} [delete]
func DeletePatient(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("DELETE FROM patients WHERE id = ?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = stmt.Exec(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
