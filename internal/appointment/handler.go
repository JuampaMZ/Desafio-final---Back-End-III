package appointment

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"odontology-appointments/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

// GET: Obtener todos los turnos
// @Summary Listar todos los turnos
// @Tags Turno
// @Produce json
// @Success 200 {array} models.Appointment
// @Router /turnos [get]
func GetAllAppointments(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM appointments")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var appointments []models.Appointment
		for rows.Next() {
			var appointment models.Appointment
			rows.Scan(&appointment.ID, &appointment.Date, &appointment.Time, &appointment.Description, &appointment.PatientID, &appointment.DentistID)
			appointments = append(appointments, appointment)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(appointments)
	}
}

// POST: Crear un nuevo turno
// @Summary Agregar un nuevo turno
// @Tags Turno
// @Accept json
// @Produce json
// @Param turno body models.Appointment true "Turno"
// @Success 201 {object} models.Appointment
// @Router /turnos [post]
func CreateAppointment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appointment models.Appointment
		err := json.NewDecoder(r.Body).Decode(&appointment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("INSERT INTO appointments (date, time, description, patient_id, dentist_id) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res, err := stmt.Exec(appointment.Date, appointment.Time, appointment.Description, appointment.PatientID, appointment.DentistID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, _ := res.LastInsertId()
		appointment.ID = int(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(appointment)
	}
}

// GET: Obtener turno por ID
// @Summary Obtener un turno por ID
// @Tags Turno
// @Produce json
// @Param id path int true "ID del turno"
// @Success 200 {object} models.Appointment
// @Failure 404 {object} models.Error
// @Router /turnos/{id} [get]
func GetAppointmentByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var appointment models.Appointment
		err = db.QueryRow("SELECT * FROM appointments WHERE id = ?", id).Scan(
			&appointment.ID, &appointment.Date, &appointment.Time, &appointment.Description, &appointment.PatientID, &appointment.DentistID)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Appointment not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(appointment)
	}
}

// PUT: Actualizar turno
// @Summary Actualizar un turno
// @Tags Turno
// @Accept json
// @Produce json
// @Param id path int true "ID del turno"
// @Param turno body models.Appointment true "Turno"
// @Success 200 {object} models.Appointment
// @Failure 404 {object} models.Error
// @Router /turnos/{id} [put]
func UpdateAppointment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var appointment models.Appointment
		err = json.NewDecoder(r.Body).Decode(&appointment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("UPDATE appointments SET date = ?, time = ?, description = ?, patient_id = ?, dentist_id = ? WHERE id = ?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = stmt.Exec(appointment.Date, appointment.Time, appointment.Description, appointment.PatientID, appointment.DentistID, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(appointment)
	}
}

// PATCH: Actualizar parcialmente turno
// @Summary Actualizar algunos campos de un turno
// @Tags Turno
// @Accept json
// @Produce json
// @Param id path int true "ID del turno"
// @Success 200 {object} models.Appointment
// @Failure 404 {object} models.Error
// @Router /turnos/{id} [patch]
func PartialUpdateAppointment(db *sql.DB) http.HandlerFunc {
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

		query := "UPDATE appointments SET "
		args := []interface{}{}
		first := true

		if date, ok := fields["date"].(string); ok {
			if !first {
				query += ", "
			}
			query += "date = ?"
			args = append(args, date)
			first = false
		}

		if time, ok := fields["time"].(string); ok {
			if !first {
				query += ", "
			}
			query += "time = ?"
			args = append(args, time)
			first = false
		}

		if description, ok := fields["description"].(string); ok {
			if !first {
				query += ", "
			}
			query += "description = ?"
			args = append(args, description)
			first = false
		}

		if patientID, ok := fields["patient_id"].(float64); ok {
			if !first {
				query += ", "
			}
			query += "patient_id = ?"
			args = append(args, int(patientID))
			first = false
		}

		if dentistID, ok := fields["dentist_id"].(float64); ok {
			if !first {
				query += ", "
			}
			query += "dentist_id = ?"
			args = append(args, int(dentistID))
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

// DELETE: Eliminar turno
// @Summary Eliminar un turno
// @Tags Turno
// @Produce json
// @Param id path int true "ID del turno"
// @Success 204
// @Failure 404 {object} models.Error
// @Router /turnos/{id} [delete]
func DeleteAppointment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("DELETE FROM appointments WHERE id = ?")
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
