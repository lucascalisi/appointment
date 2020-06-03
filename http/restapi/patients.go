package restapi

import (
	"github.com/appointment/http/models"
	"github.com/appointment/http/restapi/operations/patients"
	rec "github.com/appointment/resources"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type patientGetter interface {
	GetPatientByID(patientID int64) (rec.Patient, error)
}

func getPatient(stg patientGetter) patients.GetPatientbyIDHandlerFunc {
	return func(params patients.GetPatientbyIDParams) middleware.Responder {
		result, err := stg.GetPatientByID(params.ID)
		if err != nil {
			return patients.NewGetPatientbyIDInternalServerError().WithPayload(newRestApiError(err))
		}

		birthDay := strfmt.Date(result.BirthDay)
		return patients.NewGetPatientbyIDOK().WithPayload(&models.Patient{
			ID:       result.ID,
			Dni:      &result.DNI,
			Name:     &result.Name,
			Sex:      &result.Sex,
			BirthDay: &birthDay,
		})
	}
}

type patientAppointmentGetter interface {
	GetPatientAppointments(patientID int64) ([]rec.Appointment, error)
}

func getPatientAppointments(stg patientAppointmentGetter) patients.GetAppointmentsByPatientHandlerFunc {
	return func(params patients.GetAppointmentsByPatientParams) middleware.Responder {
		appointmentsSearched, err := stg.GetPatientAppointments(params.ID)

		if err != nil {
			return patients.NewGetAppointmentsByPatientInternalServerError().WithPayload(newRestApiError(err))
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

		return patients.NewGetAppointmentsByPatientOK().WithPayload(result)
	}
}
