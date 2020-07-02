package restapi

import (
	"fmt"
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
			Status:         params.Status,
		})

		if err != nil {
			fmt.Println(err)
			return appointments.NewSearchAppointmentInternalServerError().WithPayload(newRestApiError(err))
		}

		result := []*models.Appointment{}
		for _, a := range appointmentsSearched {
			thisAppointment := a
			thisPatient := a.Patient
			thisProfessional := a.Professional
			thisSpecialty := a.Specialty

			date := strfmt.DateTime(thisAppointment.Date)

			professional := dbProfessionalToModelProfessional(thisProfessional)
			patient := dbPatientToModelPatient(thisPatient)

			specialty := models.Specialty{
				ID:            thisSpecialty.ID,
				Category:      &thisSpecialty.Category,
				IDSubcategory: thisSpecialty.SubCategories[0].ID,
				SubCategory:   thisSpecialty.SubCategories[0].SubCategory,
			}

			appointment := models.Appointment{
				ID:           thisAppointment.ID,
				Date:         &date,
				Status:       models.AppointmentStatus(thisAppointment.Status),
				Patient:      patient,
				Professional: professional,
				Specialty:    &specialty,
			}
			result = append(result, &appointment)
		}

		return appointments.NewSearchAppointmentOK().WithPayload(result)
	}
}
