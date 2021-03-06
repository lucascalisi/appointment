package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	rec "github.com/appointment/resources"
)

func (db *DB) SearchAppointment(params rec.AppointmentSearch) ([]rec.Appointment, error) {
	statusFilter := strings.Join(params.Status, "','")
	query := fmt.Sprintf(`SELECT id, idProfessional, status, date, idPatient, idSpecialityDetail
	FROM appointments
	WHERE (idProfessional = ?
		OR ? = 0)
	AND (idSpecialityDetail = ?
		OR ? = 0)
	AND (idPatient = ?
		OR ? = 0)
	AND (date > ?
		OR ? = '')
	AND (date < ?
		OR ? = '')
	AND status IN ('%s')`, statusFilter)

	fmt.Println(params.IDSpecialty)
	rows, err := db.Query(query, params.IDProfessional, params.IDProfessional, params.IDSpecialty, params.IDSpecialty, params.IDPatient, params.IDPatient, params.StartDate, params.StartDate, params.FinishDate, params.FinishDate)
	if err != nil {
		return nil, rec.NewStorageError(fmt.Sprintf("could not get user info  : %v", err))
	}

	defer rows.Close()

	result := []rec.Appointment{}
	var idProfessional *int64
	var idPatient *int64
	var idSpecialityDetail *int64
	for rows.Next() {
		a := rec.Appointment{}
		err := rows.Scan(&a.ID, &idProfessional, &a.Status, &a.Date, &idPatient, &idSpecialityDetail)
		if err != nil {
			fmt.Println(err)
			return nil, rec.StorageError{
				Description: fmt.Sprintf("could not search appointments: %v", err),
			}
		}

		if idProfessional != nil {
			a.Professional, err = db.GetProfessionalByID(*idProfessional)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
		}

		if idPatient != nil {
			a.Patient, _ = db.GetPatientByID(*idPatient)
		}

		if idSpecialityDetail != nil {
			a.Specialty, _ = db.GetSpecialtyByDetail(*idSpecialityDetail)
		}
		result = append(result, a)
	}

	return result, nil
}

func (db *DB) GetAppointmentById(id int64) (rec.Appointment, error) {
	a := rec.Appointment{}
	var idProfessional *int64
	var idPatient *int64
	var idSpecialityDetail *int64
	err := db.QueryRow("SELECT id, idProfessional, status, date, idPatient, idSpecialityDetail FROM appointments WHERE id = ?", id).Scan(&a.ID, &idProfessional, &a.Status, &a.Date, &idPatient, &idSpecialityDetail)
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
		a.Patient, err = db.GetPatientByID(*idPatient)
		if err != nil {
			return rec.Appointment{}, err
		}
	}

	if idSpecialityDetail != nil {
		a.Specialty, err = db.GetSpecialtyByDetail(*idSpecialityDetail)
		if err != nil {
			return rec.Appointment{}, err
		}
	}

	return a, nil
}

func (db *DB) CreateAppointment(a rec.Appointment) error {
	_, err := db.Exec("INSERT INTO appointments (idProfessional, status, date, idPatient, idSpecialityDetail) VALUES(?, ?, ?, ?, ?)", a.Professional.ID, "avaiable", a.Date, nil, a.Specialty.SubCategories[0].ID)
	if err != nil {
		return rec.NewStorageError(fmt.Sprintf("could not insert appoint: %v", err))
	}

	return nil
}
