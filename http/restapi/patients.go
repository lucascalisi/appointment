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

		return patients.NewGetPatientbyIDOK().WithPayload(dbPatientToModelPatient(result))
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

		return patients.NewGetAppointmentsByPatientOK().WithPayload(result)
	}
}

type patientAppointmentRequester interface {
	RequestAppointment(patientID int64, appointmentID int64) error
	GetAppointmentById(id int64) (rec.Appointment, error)
	GetPatientAppointments(patientID int64) ([]rec.Appointment, error)
	GetPatientByID(patientID int64) (rec.Patient, error)
}

func patientRequestAppointment(stg patientAppointmentRequester, appointmentDuration int) patients.RequestAppointmentForPatientHandlerFunc {
	return func(params patients.RequestAppointmentForPatientParams) middleware.Responder {
		patient, err := stg.GetPatientByID(params.ID)
		if err != nil {
			return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(err))
		}
		if !patient.PaymentStatus {
			return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(rec.PatientNotPaymentOK))
		}

		appointment, err := stg.GetAppointmentById(params.IDAppointment)
		if err != nil {
			return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(err))
		}

		if time.Now().After(appointment.Date) {
			return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(rec.OutDatedAppointment))
		}

		//check if requested appointment is avaiable
		if appointment.Status != "avaiable" {
			return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(rec.AppointmentNotAvaiable))
		}

		if appointment.Professional.ID == patient.ID {
			return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(rec.ProfessionalSamePatient))
		}

		appointments, err := stg.GetPatientAppointments(params.ID)
		if err != nil {
			return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(err))
		}

		addAppointmentDuration := time.Duration(appointmentDuration) * time.Minute
		for _, a := range appointments {
			thisAppointment := a
			//check if patient request two appointments for the same specialty in the same day
			if thisAppointment.Date.Day() == appointment.Date.Day() && thisAppointment.Date.Month() == appointment.Date.Month() && thisAppointment.Date.Year() == appointment.Date.Year() {
				if thisAppointment.Specialty.SubCategories[0].ID == appointment.Specialty.SubCategories[0].ID {
					return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(rec.MoreThanSpecialtyAppointment))
				}
			}
			//check overlap appointments for patient
			thisAppFinishTime := time.Time(thisAppointment.Date).Add(addAppointmentDuration)
			requestedAppointmentFinishTime := time.Time(appointment.Date).Add(addAppointmentDuration)
			if (thisAppointment.Date.After(requestedAppointmentFinishTime) && appointment.Date.After(thisAppFinishTime)) || (thisAppointment.Date == appointment.Date) {
				return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(rec.OverlapAppointment))
			}
		}

		err = stg.RequestAppointment(params.ID, params.IDAppointment)
		if err != nil {
			return patients.NewRequestAppointmentForPatientInternalServerError().WithPayload(newRestApiError(err))
		}
		return patients.NewRequestAppointmentForPatientOK()
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

func dbPatientToModelPatient(patient rec.Patient) *models.Patient {
	birthDay := strfmt.Date(patient.BirthDay)
	return &models.Patient{
		ID:            patient.ID,
		Dni:           &patient.DNI,
		Name:          &patient.Name,
		Sex:           &patient.Sex,
		BirthDay:      &birthDay,
		PaymentStatus: patient.PaymentStatus,
	}
}
