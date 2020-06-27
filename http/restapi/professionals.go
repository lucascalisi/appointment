package restapi

import (
	"fmt"
	"time"

	"github.com/appointment/http/models"
	"github.com/appointment/http/restapi/operations/professionals"
	rec "github.com/appointment/resources"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type professionalGetter interface {
	GetProfessionalByID(professionalID int64) (rec.Professional, error)
}

func getProfessional(stg professionalGetter) professionals.GetProfessionalbyIDHandlerFunc {
	return func(params professionals.GetProfessionalbyIDParams) middleware.Responder {
		result, err := stg.GetProfessionalByID(params.ID)
		if err != nil {
			return professionals.NewGetProfessionalbyIDInternalServerError().WithPayload(newRestApiError(err))
		}
		return professionals.NewGetProfessionalbyIDOK().WithPayload(dbProfessionalToModelProfessional(result))
	}
}

type professionalAppointmentsGetter interface {
	GetProfessionalAppointments(professionalID int64) ([]rec.Appointment, error)
}

func getProfessionalAppointments(stg professionalAppointmentsGetter) professionals.GetAppointmentsByprofessionalHandlerFunc {
	return func(params professionals.GetAppointmentsByprofessionalParams) middleware.Responder {
		appointmentsSearched, err := stg.GetProfessionalAppointments(params.ID)

		if err != nil {
			return professionals.NewGetAppointmentsByprofessionalInternalServerError().WithPayload(newRestApiError(err))
		}

		result := []*models.Appointment{}
		for _, a := range appointmentsSearched {
			thisAppointment := a
			thisPatient := a.Patient
			thisProfessional := a.Professional

			date := strfmt.DateTime(thisAppointment.Date)

			professional := dbProfessionalToModelProfessional(thisProfessional)
			patient := dbPatientToModelPatient(thisPatient)

			appointment := models.Appointment{
				ID:           thisAppointment.ID,
				Date:         &date,
				Status:       models.AppointmentStatus(thisAppointment.Status),
				Patient:      patient,
				Professional: professional,
			}
			result = append(result, &appointment)
		}

		return professionals.NewGetAppointmentsByprofessionalOK().WithPayload(result)
	}
}

type professionalAppointmentGetter interface {
	GetProfessionalAppointment(professionalID int64, appointmentID int64) (rec.Appointment, error)
}

func getProfessionalAppointment(stg professionalAppointmentGetter) professionals.GetAppointmentByProfessionalAppointmentIDHandlerFunc {
	return func(params professionals.GetAppointmentByProfessionalAppointmentIDParams) middleware.Responder {
		appointmentsSearched, err := stg.GetProfessionalAppointment(params.ID, params.Idappointment)

		if err != nil {
			return professionals.NewGetAppointmentByProfessionalAppointmentIDInternalServerError().WithPayload(newRestApiError(err))
		}

		result := models.Appointment{}

		thisAppointment := appointmentsSearched
		thisPatient := thisAppointment.Patient
		thisProfessional := thisAppointment.Professional

		date := strfmt.DateTime(thisAppointment.Date)
		professional := dbProfessionalToModelProfessional(thisProfessional)
		patient := dbPatientToModelPatient(thisPatient)

		result = models.Appointment{
			ID:           thisAppointment.ID,
			Date:         &date,
			Status:       models.AppointmentStatus(thisAppointment.Status),
			Patient:      patient,
			Professional: professional,
		}

		return professionals.NewGetAppointmentByProfessionalAppointmentIDOK().WithPayload(&result)
	}
}

type professionalSpecialtiessGetter interface {
	GetProfessionalSpecialties(professionalID int64) ([]rec.Specialty, error)
}

func getProfessionalSpecialties(stg professionalSpecialtiessGetter) professionals.GetSpecialtiesByProfessionalHandlerFunc {
	return func(params professionals.GetSpecialtiesByProfessionalParams) middleware.Responder {
		specialtiesGetted, err := stg.GetProfessionalSpecialties(params.ID)
		if err != nil {
			return professionals.NewGetSpecialtiesByProfessionalInternalServerError().WithPayload(newRestApiError(err))
		}

		result := []*models.Specialty{}
		for _, s := range specialtiesGetted {
			thisSpecialty := s
			for _, sd := range thisSpecialty.SubCategories {
				thisSpecialtyDetail := sd
				var subCategory string = ""
				if thisSpecialty.Category != thisSpecialtyDetail.SubCategory {
					subCategory = thisSpecialtyDetail.SubCategory
				}
				result = append(result, &models.Specialty{
					ID:            thisSpecialty.ID,
					Category:      &thisSpecialty.Category,
					IDSubcategory: thisSpecialtyDetail.ID,
					SubCategory:   subCategory,
				})
			}
		}
		return professionals.NewGetSpecialtiesByProfessionalOK().WithPayload(result)
	}
}

type professionalSchedulesGetter interface {
	GetProfessionalSchedule(professionalID int64, specialtyID int64) ([]rec.Scheduler, error)
}

func getProfessionalSchedules(stg professionalSchedulesGetter) professionals.GetProfesionalScheduleBySpecialtyHandlerFunc {
	return func(params professionals.GetProfesionalScheduleBySpecialtyParams) middleware.Responder {
		schedulesGetted, err := stg.GetProfessionalSchedule(params.ID, params.IDSpecialty)
		if err != nil {
			return professionals.NewGetProfesionalScheduleBySpecialtyInternalServerError().WithPayload(newRestApiError(err))
		}

		result := []*models.Schedule{}
		for _, s := range schedulesGetted {
			thisSchedule := s
			items := []*models.ScheduleItems0{}
			for _, si := range s.Schedule {
				thisScheduleItem := si
				items = append(items, &models.ScheduleItems0{
					Day:        thisScheduleItem.Day,
					FinishTime: thisScheduleItem.FinishTime,
					StartTime:  thisScheduleItem.StartTime,
				})
			}
			result = append(result, &models.Schedule{
				ID:       thisSchedule.ID,
				Month:    &thisSchedule.Month,
				Year:     &thisSchedule.Year,
				Schedule: items,
			})

		}
		return professionals.NewGetProfesionalScheduleBySpecialtyOK().WithPayload(result)
	}
}

type professionalSchedulesSetter interface {
	SetProfessionalSchedule(professionalID int64, specialtyID int64, scheduler rec.Scheduler) (rec.Scheduler, error)
	GetProfessionalSpecialty(professionalID int64, IDSpecialty int64) (rec.Specialty, error)
	GetProfessionalSchedule(professionalID int64, specialtyID int64) ([]rec.Scheduler, error)
	CreateAppointment(rec.Appointment) error
}

func setProfessionalSchedules(stg professionalSchedulesSetter, appointmentDuration int) professionals.SetProfesionalScheduleBySpecialtyHandlerFunc {
	return func(params professionals.SetProfesionalScheduleBySpecialtyParams) middleware.Responder {
		_, err := stg.GetProfessionalSpecialty(params.ID, params.IDSpecialty)
		if err != nil {
			return professionals.NewSetProfesionalScheduleBySpecialtyInternalServerError().WithPayload(newRestApiError(err))
		}

		paramSchedule := params.Schedule
		currentDate := time.Now()
		layout := "2006-01-02T15:04:05"
		if *paramSchedule.Month < 10 {
			layout = "2006-1-02T15:04:05"
		}

		str := fmt.Sprintf("%d-%d-01T00:00:00", *paramSchedule.Year, *paramSchedule.Month)
		scheduleTime, err := time.Parse(layout, str)
		if err != nil {
			return professionals.NewSetProfesionalScheduleBySpecialtyInternalServerError().WithPayload(newRestApiError(err))
		}
		if currentDate.After(scheduleTime) {
			return professionals.NewSetProfesionalScheduleBySpecialtyInternalServerError().WithPayload(newRestApiError(rec.EditBeforeDate))
		}

		year, month, _, _, _, _ := diffDate(scheduleTime, currentDate)
		if year > 1 || month > 2 || month < 0 {
			return professionals.NewSetProfesionalScheduleBySpecialtyInternalServerError().WithPayload(newRestApiError(rec.EditMoreThanTwoMonths))
		}
		schedule := modelScheduleToDBSchedule(paramSchedule)
		newSchedule, err := stg.SetProfessionalSchedule(params.ID, params.IDSpecialty, schedule)
		if err != nil {
			fmt.Println(err)
			return professionals.NewSetProfesionalScheduleBySpecialtyInternalServerError().WithPayload(newRestApiError(err))
		}

		appointments, err := generateAppointmentsbySchedule(params.ID, params.IDSpecialty, newSchedule, appointmentDuration)
		if err != nil {
			return professionals.NewSetProfesionalScheduleBySpecialtyInternalServerError().WithPayload(newRestApiError(err))
		}

		for _, a := range appointments {
			err := stg.CreateAppointment(a)
			if err != nil {
				return professionals.NewSetProfesionalScheduleBySpecialtyInternalServerError().WithPayload(newRestApiError(err))
			}

		}

		return professionals.NewSetProfesionalScheduleBySpecialtyOK()
	}
}

func generateAppointmentsbySchedule(idProfesional int64, idSpecialty int64, sch rec.Scheduler, appointmentDuration int) ([]rec.Appointment, error) {
	appointments := []rec.Appointment{}
	for _, item := range sch.Schedule {
		layout := "2006-01-02T15:04:05"
		startDate := fmt.Sprintf("2006-01-02T%v:00", item.StartTime)
		finishDate := fmt.Sprintf("2006-01-02T%v:00", item.FinishTime)
		scheduleStartTime, err := time.Parse(layout, startDate)
		if err != nil {
			return nil, rec.NewServiceError(fmt.Sprintf("could not parse startDate"))
		}
		scheduleFinishTime, err := time.Parse(layout, finishDate)
		if err != nil {
			return nil, rec.NewServiceError(fmt.Sprintf("could not parse finishDate"))
		}
		diff := scheduleFinishTime.Sub(scheduleStartTime)

		cantAppointments := int(diff.Minutes()) / appointmentDuration
		layout = "2006-01-02"
		if sch.Month < 10 {
			layout = "2006-1-02"
		}
		firstDayOfMonth := fmt.Sprintf("%d-%d-01", sch.Year, sch.Month)
		start, err := time.Parse(layout, firstDayOfMonth)
		for d := start; d.Month() == start.Month(); d = d.AddDate(0, 0, 1) {
			if int64(d.Weekday()) != item.Day {
				continue
			}
			appointmentDate := time.Date(d.Year(), d.Month(), d.Day(), scheduleStartTime.Hour(), scheduleStartTime.Minute(), 0, d.Nanosecond(), d.Location())
			for i := 1; i <= cantAppointments; i++ {
				if i > 1 {
					appointmentDate = appointmentDate.Add(time.Minute * time.Duration(appointmentDuration))
				}
				appointments = append(appointments, rec.Appointment{
					Date:   appointmentDate,
					Status: "avaiable",
					Professional: rec.Professional{
						ID: idProfesional,
					},
					Specialty: rec.Specialty{
						ID:            0,
						SubCategories: []rec.SpecialtyDetails{{ID: idSpecialty}},
					},
				})
			}
		}
	}
	return appointments, nil
}

func modelScheduleToDBSchedule(scheduler *models.Schedule) rec.Scheduler {
	schedulerItemsDB := []rec.ScheduleItems{}
	schedulerDB := rec.Scheduler{}

	for _, scheduleItem := range scheduler.Schedule {
		schedulerItemsDB = append(schedulerItemsDB, rec.ScheduleItems{
			Day:        scheduleItem.Day,
			StartTime:  scheduleItem.StartTime,
			FinishTime: scheduleItem.FinishTime,
		})
	}

	schedulerDB.Year = *scheduler.Year
	schedulerDB.Month = *scheduler.Month
	schedulerDB.Schedule = schedulerItemsDB

	return schedulerDB
}

func asd() {

	// get the current month
	year, month, _ := time.Now().Date()

	fmt.Printf("Current month: [%v]\n", month)

	// get the number of days of the current month
	t := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Total number of days in [%v], [%v] is [%v]\n", year, month, t.Day())

	// loop each day of the month
	for day := 1; day <= t.Day(); day++ {
		// do whatever you want here ...
		fmt.Println(day, month, year)
	}
}

func diffDate(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func dbProfessionalToModelProfessional(professional rec.Professional) *models.Professional {
	birthDay := strfmt.Date(professional.BirthDay)
	return &models.Professional{
		ID:           professional.ID,
		Dni:          &professional.DNI,
		Name:         &professional.Name,
		DoctorNumber: &professional.DoctorNumber,
		BirthDay:     &birthDay,
	}
}
