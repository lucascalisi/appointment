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
	case rec.AnauthorizedUser:
		return &models.Error{
			Description: err.Error(),
			Code:        401,
			Type:        "access denied",
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