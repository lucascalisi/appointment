package restapi

import (
	"github.com/appointment/http/models"
	"github.com/appointment/http/restapi/operations/login"
	rec "github.com/appointment/resources"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type loginAutenticator interface {
	Login(rec.LoginAuth) (rec.User, error)
}

func authenticate(stg loginAutenticator) login.LoginHandlerFunc {
	return func(params login.LoginParams) middleware.Responder {
		user, err := stg.Login(rec.LoginAuth{
			User:     *params.Credentials.User,
			Password: *params.Credentials.Password,
		})

		if err != nil {
			return login.NewLoginInternalServerError().WithPayload(newRestApiError(err))
		}

		roles := []*models.Role{}
		for _, r := range user.Roles {
			thisRole := r
			roles = append(roles, &models.Role{
				ID:          thisRole.ID,
				Name:        &thisRole.Name,
				Description: &thisRole.Description,
			})
		}

		return login.NewLoginOK().WithPayload(&models.User{
			ID:     user.ID,
			Email:  strfmt.Email(user.Email),
			Roles:  roles,
			Status: user.Status,
		})
	}
}
