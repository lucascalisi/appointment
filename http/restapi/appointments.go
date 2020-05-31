package restapi

import (
	"time"

	"github.com/appointment/http/models"
	"github.com/appointment/http/restapi/operations/appointments"
	rec "github.com/appointment/resources"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type appointmentSearcher interface {
	SearchAppointment(rec.AppointmentSearch) ([]rec.Appointment, error)
}

func searchAppointments(stg appointmentSearcher) appointments.SearchAppointmentHandlerFunc {
	return func(params appointments.SearchAppointmentParams) middleware.Responder {
		timezone, _ := time.LoadLocation("America/Buenos_Aires")
		appointmentsSearched, err := stg.SearchAppointment(rec.AppointmentSearch{
			IDPatient:      *params.IDPatient,
			IDProfessional: *params.IDProfessional,
			StartDate:      time.Time(*params.StartDate).In(timezone).Add(3 * time.Hour),
			FinishDate:     time.Time(*params.FinishDate).In(timezone).Add(3 * time.Hour),
			IDSpecialty:    *params.Idspecialty,
			Status:         *params.Status,
		})

		if err != nil {
			return appointments.NewSearchAppointmentInternalServerError().WithPayload(newRestApiError(err))
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

		return appointments.NewSearchAppointmentOK().WithPayload(result)
	}
}