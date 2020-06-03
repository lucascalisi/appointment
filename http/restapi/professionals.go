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

		birthDay := strfmt.Date(result.BirthDay)
		return professionals.NewGetProfessionalbyIDOK().WithPayload(&models.Professional{
			ID:           result.ID,
			Dni:          &result.DNI,
			Name:         &result.Name,
			DoctorNumber: &result.DoctorNumber,
			BirthDay:     &birthDay,
		})
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
			birthDayPatient := strfmt.Date(thisPatient.BirthDay)
			birthDayProfessional := strfmt.Date(thisProfessional.BirthDay)

			professional := models.Professional{
				ID:           thisProfessional.ID,
				Dni:          &thisProfessional.DNI,
				Name:         &thisProfessional.Name,
				DoctorNumber: &thisProfessional.DoctorNumber,
				BirthDay:     &birthDayProfessional,
			}

			patient := models.Patient{
				ID:       thisPatient.ID,
				Dni:      &thisPatient.DNI,
				Name:     &thisPatient.Name,
				Sex:      &thisPatient.Sex,
				BirthDay: &birthDayPatient,
			}

			appointment := models.Appointment{
				ID:           thisAppointment.ID,
				Date:         &date,
				Status:       models.AppointmentStatus(thisAppointment.Status),
				Patient:      &patient,
				Professional: &professional,
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
		birthDayPatient := strfmt.Date(thisPatient.BirthDay)
		birthDayProfessional := strfmt.Date(thisProfessional.BirthDay)

		professional := models.Professional{
			ID:           thisProfessional.ID,
			Dni:          &thisProfessional.DNI,
			Name:         &thisProfessional.Name,
			DoctorNumber: &thisProfessional.DoctorNumber,
			BirthDay:     &birthDayProfessional,
		}

		patient := models.Patient{
			ID:       thisPatient.ID,
			Dni:      &thisPatient.DNI,
			Name:     &thisPatient.Name,
			Sex:      &thisPatient.Sex,
			BirthDay: &birthDayPatient,
		}

		result = models.Appointment{
			ID:           thisAppointment.ID,
			Date:         &date,
			Status:       models.AppointmentStatus(thisAppointment.Status),
			Patient:      &patient,
			Professional: &professional,
		}

		return professionals.NewGetAppointmentByProfessionalAppointmentIDOK().WithPayload(&result)
	}
}
