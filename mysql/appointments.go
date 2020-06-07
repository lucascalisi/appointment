package mysql

import (
	"database/sql"
	"fmt"

	rec "github.com/appointment/resources"
)

func (db *DB) SearchAppointment(params rec.AppointmentSearch) ([]rec.Appointment, error) {
	query := `SELECT id, idProfessional, status, date, idPatient
	FROM appointments
	WHERE (idProfessional = ?
		OR ? = 0)
	AND (idPatient = ?
		OR ? = 0)
	AND (date > ?
		OR ? = '')
	AND (date < ?
		OR ? = '')
	AND (status = ?
		OR ? = '')`

	rows, err := db.Query(query, params.IDProfessional, params.IDProfessional, params.IDPatient, params.IDPatient, params.StartDate, params.StartDate, params.FinishDate, params.FinishDate, params.Status, params.Status)
	if err != nil {
		return nil, rec.NewStorageError(fmt.Sprintf("could not get user info  : %v", err))
	}

	defer rows.Close()

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

		if idPatient != nil {
			a.Patient, _ = db.GetPatientByID(*idPatient)
		}
		result = append(result, a)
	}

	return result, nil
}

func (db *DB) GetAppointmentById(id int64) (rec.Appointment, error) {
	a := rec.Appointment{}
	var idProfessional *int64
	var idPatient *int64
	err := db.QueryRow("SELECT id, idProfessional, status, date, idPatient FROM appointments WHERe id = ?", id).Scan(&a.ID, &idProfessional, &a.Status, &a.Date, &idPatient)
	if err != nil {
		if err == sql.ErrNoRows {
			return rec.Appointment{}, rec.NotFound
		}
		return rec.Appointment{}, rec.NewStorageError(fmt.Sprintf("could not get user info  : %v", err))
	}

	if idProfessional != nil {
		a.Professional, err = db.GetProfessionalByID(*idProfessional)
		if err != nil {
			return rec.Appointment{}, err
		}
	}

	if idPatient != nil {
		a.Patient, _ = db.GetPatientByID(*idPatient)
	}

	return a, nil
}
