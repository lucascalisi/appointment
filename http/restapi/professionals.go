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
					ID:          thisSpecialtyDetail.ID,
					Category:    &thisSpecialty.Category,
					SubCategory: subCategory,
				})
			}
		}
		return professionals.NewGetSpecialtiesByProfessionalOK().WithPayload(result)
	}
}
