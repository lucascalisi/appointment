package restapi

import (
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
					FinishTime: rec.ScheduledTimeToString(thisScheduleItem.FinishTime),
					StartTime:  rec.ScheduledTimeToString(thisScheduleItem.StartTime),
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
