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
