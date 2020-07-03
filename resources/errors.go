package resources

import (
	"fmt"
	"runtime"
)

type ServiceError struct {
	Description string
	trace       []uintptr
}

func NewServiceError(errorstr string) ServiceError {
	return ServiceError{Description: errorstr, trace: stackTrace()}
}

func (se ServiceError) Error() string {
	return se.Description
}

func (se ServiceError) ErrorClass() string {
	return "ServiceError"
}

func (se ServiceError) StackTrace() []uintptr {
	return se.trace
}

type StorageError struct {
	Description string
	trace       []uintptr
}

func NewStorageError(errorstr string) ServiceError {
	return ServiceError{Description: errorstr, trace: stackTrace()}
}

func (se StorageError) Error() string {
	return se.Description
}

func (se StorageError) ErrorClass() string {
	return "StorageError"
}

func (se StorageError) StackTrace() []uintptr {
	return se.trace
}

func stackTrace() []uintptr {
	callers := make([]uintptr, 20)
	written := runtime.Callers(2, callers)
	return callers[0:written]
}

var UnauthorizedUser error = fmt.Errorf("user or password incorrect")
var PatientNotPaymentOK error = fmt.Errorf("patient has not payment status ok")
var AppointmentNotAvaiable error = fmt.Errorf("the requested appointment is not avaiable")
var MoreThanSpecialtyAppointment error = fmt.Errorf("could not have two appointments with the same specialty in the same day")
var OverlapAppointment error = fmt.Errorf("not could get two overlap appointments")
var OutDatedAppointment error = fmt.Errorf("not could get outdated appointment")
var ProfessionalSamePatient error = fmt.Errorf("not could get appointment with yourself!!!")
var EditBeforeDate error = fmt.Errorf("not could set scheduler before now!")
var EditMoreThanTwoMonths error = fmt.Errorf("not could edit scheduler after than two months")
var AttendOtherSpecialtyThisDay error = fmt.Errorf("not could attend more than specialty in the same day of week")
var ScheduleAlreadySetted error = fmt.Errorf("not could set again the schedule for the same year and month")
var PatientEmailNotFound error = fmt.Errorf("not could get email for patient")
var AppointmentAlreadyCancelled error = fmt.Errorf("not could cancel the appointment")
var PatientIsAlreadyWaiting error = fmt.Errorf("not could add patient to queue")
var AppointmentCancelProfessional error = fmt.Errorf("not could cancel appointment without a prior week in advance")
var NotFound error = fmt.Errorf("not found")
