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
