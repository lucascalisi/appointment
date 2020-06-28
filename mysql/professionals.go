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

func (db *DB) GetProfessionalAppointments(id int64, status string) ([]rec.Appointment, error) {
	query := `SELECT id, idProfessional, status, date, idPatient
	FROM appointments
	WHERE idProfessional = ?
	AND status = ?`

	rows, err := db.Query(query, id, status)
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

func (db *DB) GetProfessionalSpecialties(id int64) ([]rec.Specialty, error) {
	query := `SELECT s.id, s.name, sd.id, sd.name 
	FROM specialityDetailsByProfessional sdp
	INNER JOIN specialityDetails sd
		ON sdp.idSpecialityDetail = sd.id
	INNER JOIN specialties s
		ON s.id = sd.idSpeciality
	WHERE sdp.idProfessional = ?`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, rec.NewStorageError(fmt.Sprintf("could not get specialties by professional: %v", err))
	}

	defer rows.Close()

	specialties := []rec.Specialty{}
	specialtiDetails := []rec.SpecialtyDetails{}
	sNext := rec.Specialty{}
	s := rec.Specialty{}
	for rows.Next() {
		var idSubCategory *int64
		var subCategory *string

		err := rows.Scan(&sNext.ID, &sNext.Category, &idSubCategory, &subCategory)
		if err != nil {
			return nil, rec.StorageError{
				Description: fmt.Sprintf("could not search specialties: %v", err),
			}
		}

		if s.ID != sNext.ID {
			s.SubCategories = specialtiDetails
			specialties = append(specialties, s)
			specialtiDetails = nil
		}

		if idSubCategory != nil {
			specialtiDetails = append(specialtiDetails, rec.SpecialtyDetails{
				ID:          *idSubCategory,
				SubCategory: *subCategory,
			})
		}
		s = sNext
	}

	s.SubCategories = specialtiDetails
	specialties = append(specialties, s)

	return specialties, nil
}

func (db *DB) GetProfessionalSpecialty(id int64, idSpeciality int64) (rec.Specialty, error) {
	query := `SELECT s.id, s.name, sd.id, sd.name 
	FROM specialityDetailsByProfessional sdp
	INNER JOIN specialityDetails sd
		ON sdp.idSpecialityDetail = sd.id
	INNER JOIN specialties s
		ON s.id = sd.idSpeciality
	WHERE sdp.idProfessional = ?
	and sd.id = ?`

	specialty := rec.Specialty{}
	specialtyDetails := []rec.SpecialtyDetails{}
	var idSubCategory *int64
	var subCategory *string
	err := db.QueryRow(query, id, idSpeciality).Scan(&specialty.ID, &specialty.Category, &idSubCategory, &subCategory)
	if err != nil {
		return rec.Specialty{}, rec.NewStorageError(fmt.Sprintf("could not get specialty by professional: %v", err))
	}

	specialty.SubCategories = append(specialtyDetails, rec.SpecialtyDetails{
		ID:          *idSubCategory,
		SubCategory: *subCategory,
	})

	return specialty, nil
}

func (db *DB) GetProfessionalSchedule(id int64, idSpecialty int64) ([]rec.Scheduler, error) {
	query := `SELECT s.id, s.idProfessional, s.idSpeciality, s.year, s.month
	FROM professionalsScheduleBySpecialty s
	WHERE idProfessional = ?
	AND idSpeciality = ?`

	rows, err := db.Query(query, id, idSpecialty)
	if err != nil {
		return nil, rec.NewStorageError(fmt.Sprintf("could not get schedules by professional: %v", err))
	}

	defer rows.Close()

	schedulers := []rec.Scheduler{}
	for rows.Next() {
		s := rec.Scheduler{}
		var idProfessional *int64
		err := rows.Scan(&s.ID, &idProfessional, &s.IDSpecialty, &s.Year, &s.Month)
		if err != nil {
			return nil, rec.StorageError{
				Description: fmt.Sprintf("could not get schedules: %v", err),
			}
		}

		s.Schedule, err = db.GetSchedulerItemsByID(s.ID)
		if err != nil {
			return nil, rec.StorageError{
				Description: fmt.Sprintf("could not get schedule items: %v", err),
			}
		}
		schedulers = append(schedulers, s)
	}

	return schedulers, nil
}

func (db *DB) SearchSchedule(id int64, year int64, month int64) ([]rec.Scheduler, error) {
	query := `SELECT s.id, s.idProfessional, s.idSpeciality, s.year, s.month
	FROM professionalsScheduleBySpecialty s
	WHERE idProfessional = ?
	AND year = ?
	AND month = ?`

	rows, err := db.Query(query, id, year, month)
	if err != nil {
		return nil, rec.NewStorageError(fmt.Sprintf("could not get schedules by professional: %v", err))
	}

	defer rows.Close()

	schedulers := []rec.Scheduler{}
	for rows.Next() {
		s := rec.Scheduler{}
		var idProfessional *int64
		err := rows.Scan(&s.ID, &idProfessional, &s.IDSpecialty, &s.Year, &s.Month)
		if err != nil {
			return nil, rec.StorageError{
				Description: fmt.Sprintf("could not get schedules: %v", err),
			}
		}

		s.Schedule, err = db.GetSchedulerItemsByID(s.ID)
		if err != nil {
			return nil, rec.StorageError{
				Description: fmt.Sprintf("could not get schedule items: %v", err),
			}
		}
		schedulers = append(schedulers, s)
	}

	return schedulers, nil
}

func (db *DB) SetProfessionalSchedule(id int64, idSpecialty int64, scheduler rec.Scheduler) (rec.Scheduler, error) {
	result, err := db.Exec("INSERT INTO professionalsScheduleBySpecialty (idProfessional, idSpeciality, year, month) VALUES(?, ?, ?, ?)", id, idSpecialty, scheduler.Year, scheduler.Month)
	if err != nil {
		return rec.Scheduler{}, rec.NewStorageError(fmt.Sprintf("could not insert schedule: %v", err))
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return rec.Scheduler{}, rec.NewStorageError(fmt.Sprintf("could not get inserted schedulerid: %v", err))
	}

	for _, schedulerItem := range scheduler.Schedule {
		_, err := db.Exec("INSERT INTO professionalsScheduleItemsBySpecialty (idSchedule, dayOfWeek, startTime, finishTime) VALUES(?, ?, ?, ?)", lastId, schedulerItem.Day, schedulerItem.StartTime, schedulerItem.FinishTime)
		if err != nil {
			return rec.Scheduler{}, rec.NewStorageError(fmt.Sprintf("could not insert schedule item: %v", err))
		}
	}

	schedulerItems, _ := db.GetSchedulerItemsByID(lastId)

	return rec.Scheduler{
		ID:       lastId,
		Year:     scheduler.Year,
		Month:    scheduler.Month,
		Schedule: schedulerItems,
	}, nil
}

func (db *DB) GetSchedulerItemsByID(id int64) ([]rec.ScheduleItems, error) {
	query := `SELECT s.id, s.idSchedule, s.dayOfWeek, s.startTime, s.finishTime
	FROM professionalsScheduleItemsBySpecialty s
	WHERE idSchedule = ?`

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, rec.NewStorageError(fmt.Sprintf("could not get schedule items by id scheduler: %v", err))
	}

	defer rows.Close()

	items := []rec.ScheduleItems{}
	for rows.Next() {
		item := rec.ScheduleItems{}
		var idSchedule *int64
		var startTime *string
		var finishTime *string
		err := rows.Scan(&item.ID, &idSchedule, &item.Day, &startTime, &finishTime)
		if err != nil {
			return nil, rec.StorageError{
				Description: fmt.Sprintf("could not make schedule item: %v", err),
			}
		}

		if startTime != nil {
			item.StartTime = *startTime
		}

		if finishTime != nil {
			item.FinishTime = *finishTime
		}

		items = append(items, item)
	}

	return items, nil
}

func (db *DB) GetProfessionalScheduleByID(idSchedule int64) (rec.Scheduler, error) {
	query := `SELECT s.id, s.idProfessional, s.idSpeciality, s.year, s.month
	FROM professionalsScheduleBySpecialty s
	WHERE id = ?`

	var idSpeciality *int64
	var idProfessional *int64
	s := rec.Scheduler{}
	err := db.QueryRow(query, idSchedule).Scan(&s.ID, &idProfessional, &idSpeciality, &s.Year, &s.Month)
	if err != nil {
		return rec.Scheduler{}, rec.NewStorageError(fmt.Sprintf("could not get schedules id by professional: %v", err))
	}

	s.Schedule, err = db.GetSchedulerItemsByID(s.ID)
	if err != nil {
		return rec.Scheduler{}, rec.StorageError{
			Description: fmt.Sprintf("could not get schedule items: %v", err),
		}
	}

	return s, nil
}
