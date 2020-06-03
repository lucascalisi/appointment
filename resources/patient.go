package resources

import "time"

type PatientStorage interface {
	getPatientByID(patientID int64) (Patient, error)
	getPatientAppointments(patientID int64) ([]Appointment, error)
}

type Patient struct {
	ID       int64
	Name     string
	DNI      int64
	Sex      string
	BirthDay time.Time
}
