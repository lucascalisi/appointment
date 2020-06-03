package mysql

import (
	"database/sql"
	"fmt"

	rec "github.com/appointment/resources"
)

func (db *DB) GetPatientByID(id int64) (rec.Patient, error) {
	result := rec.Patient{}
	err := db.QueryRow("SELECT idUser, dni, name, sex, birthDay FROM patients WHERe idUser = ?", id).Scan(&result.ID, &result.DNI, &result.Name, &result.Sex, &result.BirthDay)
	if err != nil {
		if err == sql.ErrNoRows {
			return rec.Patient{}, rec.NotFound
		}
		return rec.Patient{}, rec.NewStorageError(fmt.Sprintf("could not get user info  : %v", err))
	}

	return result, nil
}

func (db *DB) GetPatientAppointments(id int64) ([]rec.Appointment, error) {
	query := `SELECT id, idProfessional, status, date, idPatient
	FROM appointments
	WHERE idPatient = ?`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, rec.NewStorageError(fmt.Sprintf("could not get appointments  : %v", err))
	}

	defer rows.Close()

	patient, _ := db.GetPatientByID(id)

	result := []rec.Appointment{}
	var idProfessional *int64
	var idPatient *int64
	for rows.Next() {
		a := rec.Appointment{}
		err := rows.Scan(&a.ID, &idProfessional, &a.Status, &a.Date, &idPatient)
		if err != nil {
			return nil, rec.StorageError{
				Description: fmt.Sprintf("could not search appointments: %v", err),
			}
		}

		if idProfessional != nil {
			a.Professional, err = db.GetProfessionalByID(*idProfessional)
			if err != nil {
				return nil, err
			}
		}

		a.Patient = patient
		result = append(result, a)
	}

	return result, nil
}
