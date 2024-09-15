package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./odontology.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables(db)
	return db
}

func createTables(db *sql.DB) {
	dentistTable := `
    CREATE TABLE IF NOT EXISTS dentists (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        last_name TEXT,
        first_name TEXT,
        license TEXT
    );`

	patientTable := `
    CREATE TABLE IF NOT EXISTS patients (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        last_name TEXT,
        first_name TEXT,
        address TEXT,
        dni TEXT,
        registration_date TEXT
    );`

	appointmentTable := `
    CREATE TABLE IF NOT EXISTS appointments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        date TEXT,
        time TEXT,
        description TEXT,
        patient_id INTEGER,
        dentist_id INTEGER,
        FOREIGN KEY(patient_id) REFERENCES patients(id),
        FOREIGN KEY(dentist_id) REFERENCES dentists(id)
    );`

	if _, err := db.Exec(dentistTable); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(patientTable); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(appointmentTable); err != nil {
		log.Fatal(err)
	}
}
