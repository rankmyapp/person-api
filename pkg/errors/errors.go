package errors

import "fmt"

type ValidationError struct {
	message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s", e.message)
}

func NewValidationError(message string) error {
	return &ValidationError{message: message}
}

func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}

type ConflictError struct {
	message string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("Conflict error: %s", e.message)
}

func NewConflictError(message string) error {
	return &ConflictError{message: message}
}

func IsConflictError(err error) bool {
	_, ok := err.(*ConflictError)
	return ok
}

type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found error: %s", e.message)
}

func NewNotFoundError(message string) error {
	return &NotFoundError{message: message}
}

func IsNotFoundError(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

type InternalError struct {
	message string
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("Internal error: %s", e.message)
}

func NewInternalError(message string) error {
	return &InternalError{message: message}
}

func IsInternalError(err error) bool {
	_, ok := err.(*InternalError)
	return ok
}
