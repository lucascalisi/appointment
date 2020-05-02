package app_models

import "time"

type Patient struct {
	ID        int
	UserName  string
	Name      string
	Status    string
	DNI       int
	Sex       string
	BirthDate time.Time
}

type PatientSearch struct {
	Name string
	DNI  int
}

type PatientStorage interface {
	SearchPatient(PatientSearch) ([]Patient, error)
}
