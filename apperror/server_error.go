package apperror

type ServerError struct {
	err     error
	message string
}

func NewServerError(err error, msg ...string) *ServerError {
	message := err.Error()
	if len(msg) > 0 {
		message = msg[0]
	}
	return &ServerError{
		err:     err,
		message: message,
	}
}

func (se *ServerError) Error() string {
	return se.err.Error()
}

func (se *ServerError) Message() string {
	return se.message
}
