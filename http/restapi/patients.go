package restapi

import (
	"time"

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

type patientAppointmentRequester interface {
	RequestAppointment(patientID int64, appointmentID int64) error
	GetAppointmentById(id int64) (rec.Appointment, error)
}

func patientRequestAppointment(stg patientAppointmentRequester) patients.RequestAppointmentForPatientHandlerFunc {
	return func(params patients.RequestAppointmentForPatientParams) middleware.Responder {
		appointment, err := stg.GetAppointmentById(params.IDAppointment)
		if err != nil {
			return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(err))
		}

		if appointment.Status == "avaiable" {
			err := stg.RequestAppointment(params.ID, params.IDAppointment)
			if err != nil {
				return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(err))
			}
			return patients.NewRequestAppointmentForPatientOK()
		}

		return patients.NewRequestAppointmentForPatientInternalServerError()
	}
}

type patientAppointmentConfirmer interface {
	ConfirmAppointment(patientID int64, appointmentID int64) error
	GetAppointmentById(id int64) (rec.Appointment, error)
}

func patientConfirmAppointment(stg patientAppointmentConfirmer) patients.ConfirmAppointmentForPatientHandlerFunc {
	return func(params patients.ConfirmAppointmentForPatientParams) middleware.Responder {
		appointment, err := stg.GetAppointmentById(params.IDAppointment)
		if err != nil {
			return patients.NewConfirmAppointmentForPatientInternalServerError().WithPayload(newRestApiError(err))
		}

		if appointment.Patient.ID == params.ID && appointment.Status == "pending" {
			timeNow := time.Now()
			diff := appointment.Date.Sub(timeNow)
			if diff.Hours() > 0 && diff.Hours() < 24 {
				err := stg.ConfirmAppointment(appointment.Patient.ID, appointment.ID)
				if err != nil {
					return patients.NewConfirmAppointmentForPatientInternalServerError().WithPayload(newRestApiError(err))
				}
				return patients.NewConfirmAppointmentForPatientOK()
			}
		}

		return patients.NewConfirmAppointmentForPatientInternalServerError()
	}
}

type patientAppointmentCanceler interface {
	CancelAppointment(patientID int64, appointmentID int64) error
	GetAppointmentById(id int64) (rec.Appointment, error)
}

func patientCancelAppointment(stg patientAppointmentCanceler) patients.CancelAppoinmentForPatientHandlerFunc {
	return func(params patients.CancelAppoinmentForPatientParams) middleware.Responder {
		appointment, err := stg.GetAppointmentById(params.IDAppointment)
		if err != nil {
			return patients.NewCancelAppoinmentForPatientInternalServerError().WithPayload(newRestApiError(err))
		}

		if appointment.Patient.ID == params.ID && appointment.Status == "pending" {
			err := stg.CancelAppointment(params.ID, params.IDAppointment)
			if err != nil {
				return patients.NewCancelAppoinmentForPatientInternalServerError().WithPayload(newRestApiError(err))
			}
			return patients.NewCancelAppoinmentForPatientOK()
		}

		return patients.NewCancelAppoinmentForPatientInternalServerError()

	}
}
