package resources

import "time"

type AppointmentSearch struct {
	IDPatient      int64
	IDProfessional int64
	StartDate      time.Time
	FinishDate     time.Time
	IDSpecialty    int64
	Status         string
}

type AppointmentStorage interface {
	SearchAppointment(AppointmentSearch) ([]Appointment, error)
}

type Appointment struct {
	ID           int64
	Date         time.Time
	Status       string
	Patient      Patient
	Professional Professional
	Specialty    Specialty
}
