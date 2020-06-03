package mysql

import (
	"database/sql"
	"fmt"

	rec "github.com/appointment/resources"
)

func (db *DB) GetProfessionalByID(id int64) (rec.Professional, error) {
	result := rec.Professional{}
	err := db.QueryRow("SELECT idUser, dni, name, sex, doctorNumber, birthDay FROM professionals WHERe idUser = ?", id).Scan(&result.ID, &result.DNI, &result.Name, &result.Sex, &result.DoctorNumber, &result.BirthDay)
	if err != nil {
		if err == sql.ErrNoRows {
			return rec.Professional{}, rec.NotFound
		}
		return rec.Professional{}, rec.NewStorageError(fmt.Sprintf("could not get user info  : %v", err))
	}

	return result, nil
}

func (db *DB) GetProfessionalAppointments(id int64) ([]rec.Appointment, error) {
	query := `SELECT id, idProfessional, status, date, idPatient
	FROM appointments
	WHERE idProfessional = ?`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, rec.NewStorageError(fmt.Sprintf("could not get appointments  : %v", err))
	}

	defer rows.Close()

	professional, _ := db.GetProfessionalByID(id)

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

		if idPatient != nil {
			a.Patient, err = db.GetPatientByID(*idPatient)
			if err != nil {
				return nil, err
			}
		}

		a.Professional = professional
		result = append(result, a)
	}

	return result, nil
}

func (db *DB) GetProfessionalAppointment(id int64, idAppointment int64) (rec.Appointment, error) {
	query := `SELECT id, idProfessional, status, date, idPatient
	FROM appointments
	WHERE idProfessional = ?
	AND id = ?`

	result := rec.Appointment{}
	var idProfessional *int64
	var idPatient *int64
	err := db.QueryRow(query, id, idAppointment).Scan(&result.ID, &idProfessional, &result.Status, &result.Date, &idPatient)
	if err != nil {
		return rec.Appointment{}, rec.NewStorageError(fmt.Sprintf("could not get appointment: %v", err))
	}

	if idPatient != nil {
		result.Patient, err = db.GetPatientByID(*idPatient)
		if err != nil {
			return rec.Appointment{}, err
		}
	}

	if idProfessional != nil {
		result.Professional, err = db.GetProfessionalByID(*idProfessional)
		if err != nil {
			return rec.Appointment{}, err
		}
	}
	return result, nil
}
