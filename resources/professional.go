package resources

import "time"

type ProfessionalStorage interface {
	getProfessionalByID(professionalID int64) (Professional, error)
}

type Professional struct {
	ID           int64
	Name         string
	DNI          int64
	Sex          string
	DoctorNumber int64
	BirthDay     time.Time
}
