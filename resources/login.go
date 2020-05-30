package resources

import "github.com/go-openapi/strfmt"

type LoginAuth struct {
	User     string
	Password strfmt.Password
}

type LoginStorage interface {
	Login(LoginAuth) error
}
