package restapi

import (
	"github.com/appointment/http/restapi/operations/login"
	rec "github.com/appointment/resources"
	middleware "github.com/go-openapi/runtime/middleware"
)

type loginAutenticator interface {
	Login(rec.LoginAuth) error
}

func authenticate(stg loginAutenticator) login.LoginHandlerFunc {
	return func(params login.LoginParams) middleware.Responder {
		err := stg.Login(rec.LoginAuth{
			User:     *params.Credentials.User,
			Password: *params.Credentials.Password,
		})

		if err != nil {
			return login.NewLoginInternalServerError().WithPayload(newRestApiError(err))
		}
		return login.NewLoginOK()
	}
}
