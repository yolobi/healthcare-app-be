package apperror

type ValidationError struct {
	err     error
	message string
}

func NewValidationError(err error, msg ...string) *ValidationError {
	message := err.Error()
	if len(msg) > 0 {
		message = msg[0]
	}
	return &ValidationError{
		err:     err,
		message: message,
	}
}

func (se *ValidationError) Error() string {
	return se.err.Error()
}

func (se *ValidationError) Message() string {
	return se.message
}
