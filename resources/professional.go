package resources

import "time"

type ProfessionalStorage interface {
	getProfessionalByID(professionalID int64) (Professional, error)
	getProfessionalAppointments(professionalID int64, status string) ([]Appointment, error)
	getProfessionalAppointment(professionalID int64, appointmentID int64) (Appointment, error)
	getProfessionalSpecialties(professionalID int64) ([]Specialty, error)
	getProfessionalSchedule(professionalID int64, specialtyID int64) ([]Scheduler, error)
	setProfessionalSchedule(professionalID int64, specialtyID int64, scheduler Scheduler) (Scheduler, error)
	getProfessionalSpecialty(professionalID int64, IDSpecialty int64) (bool, error)
	searchSchedule(id int64, year int64, month int64) ([]Scheduler, error)
}

type Professional struct {
	ID           int64
	Name         string
	DNI          int64
	Sex          string
	DoctorNumber int64
	BirthDay     time.Time
}
