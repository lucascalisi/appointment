package app_models

import (
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
