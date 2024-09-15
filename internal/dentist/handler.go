package dentist

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"odontology-appointments/pkg/models"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Listar todos los dentistas
// @Tags Dentista
// @Produce json
// @Success 200 {array} models.Dentist
// @Router /dentistas [get]
func GetAllDentists(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Consulta a la base de datos para obtener todos los dentistas
		rows, err := db.Query("SELECT id, last_name, first_name, license FROM dentists")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Crear una lista para almacenar los dentistas
		var dentists []models.Dentist

		// Iterar sobre los resultados y agregar cada dentista a la lista
		for rows.Next() {
			var dentist models.Dentist
			if err := rows.Scan(&dentist.ID, &dentist.LastName, &dentist.FirstName, &dentist.License); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dentists = append(dentists, dentist)
		}

		// Verificar si hubo un error al iterar
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Configurar el encabezado de la respuesta y codificar la lista de dentistas como JSON
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dentists); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// POST: Agregar dentista
// @Summary Agregar un nuevo dentista
// @Tags Dentista
// @Accept json
// @Produce json
// @Param dentista body models.Dentist true "Dentista"
// @Success 201 {object} models.Dentist
// @Failure 404 {object} models.Error
// @Router /dentistas [post]
func CreateDentist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dentist models.Dentist
		err := json.NewDecoder(r.Body).Decode(&dentist)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("INSERT INTO dentists (last_name, first_name, license) VALUES (?, ?, ?)")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res, err := stmt.Exec(dentist.LastName, dentist.FirstName, dentist.License)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, _ := res.LastInsertId()
		dentist.ID = int(id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dentist)
	}
}

// GET: Traer dentista por ID
// @Summary Obtener un dentista por ID
// @Tags Dentista
// @Produce json
// @Param id path int true "ID del dentista"
// @Success 200 {object} models.Dentist
// @Failure 404 {object} models.Error
// @Router /dentistas/{id} [get]
func GetDentistByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var dentist models.Dentist
		err = db.QueryRow("SELECT * FROM dentists WHERE id = ?", id).Scan(
			&dentist.ID, &dentist.LastName, &dentist.FirstName, &dentist.License)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Dentist not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dentist)
	}
}

// PUT: Actualizar dentista
// @Summary Actualizar un dentista
// @Tags Dentista
// @Accept json
// @Produce json
// @Param id path int true "ID del dentista"
// @Param dentista body models.Dentist true "Dentista"
// @Success 200 {object} models.Dentist
// @Failure 404 {object} models.Error
// @Router /dentistas/{id} [put]
func UpdateDentist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		var dentist models.Dentist
		err = json.NewDecoder(r.Body).Decode(&dentist)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("UPDATE dentists SET last_name = ?, first_name = ?, license = ? WHERE id = ?")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = stmt.Exec(dentist.LastName, dentist.FirstName, dentist.License, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dentist)
	}
}

// PATCH: Actualizar dentista por algún campo
// @Summary Actualizar algunos campos de un dentista
// @Tags Dentista
// @Accept json
// @Produce json
// @Param id path int true "ID del dentista"
// @Success 200 {object} models.Dentist
// @Failure 404 {object} models.Error
// @Router /dentistas/{id} [patch]
func PartialUpdateDentist(db *sql.DB) http.HandlerFunc {
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

		query := "UPDATE dentists SET "
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

		if license, ok := fields["license"].(string); ok {
			if !first {
				query += ", "
			}
			query += "license = ?"
			args = append(args, license)
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

// DELETE: Eliminar dentista
// @Summary Eliminar un dentista
// @Tags Dentista
// @Produce json
// @Param id path int true "ID del dentista"
// @Success 204
// @Failure 404 {object} models.Error
// @Router /dentistas/{id} [delete]
func DeleteDentist(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		stmt, err := db.Prepare("DELETE FROM dentists WHERE id = ?")
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
