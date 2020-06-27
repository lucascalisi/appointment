package restapi

import (
	"github.com/appointment/http/models"
	rec "github.com/appointment/resources"
)

func newRestApiError(err error) *models.Error {
	result := &models.Error{
		Description: err.Error(),
	}

	switch err {
	case rec.UnauthorizedUser:
		return &models.Error{
			Description: err.Error(),
			Code:        401,
			Type:        "access denied",
		}
	case rec.NotFound:
		return &models.Error{
			Description: err.Error(),
			Code:        404,
			Type:        "not found",
		}
	case rec.PatientNotPaymentOK:
		return &models.Error{
			Description: err.Error(),
			Code:        403,
			Type:        "patient has not payment ok",
		}
	case rec.AppointmentNotAvaiable:
		return &models.Error{
			Description: err.Error(),
			Code:        403,
			Type:        "appointment not avaiable",
		}
	case rec.MoreThanSpecialtyAppointment:
		return &models.Error{
			Description: err.Error(),
			Code:        403,
			Type:        "could not have two appointments with the same specialty in the same day",
		}
	case rec.OverlapAppointment:
		return &models.Error{
			Description: err.Error(),
			Code:        403,
			Type:        "overlap appointment",
		}
	}
	switch err.(type) {
	default:
		result.Code = 500
		result.Type = "internal server error"
		result.Description = "internal server error"
	}

	return result
}
