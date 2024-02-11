package apperror

type ClientError struct {
	err     error
	message string
}

func NewClientError(err error, msg ...string) *ClientError {
	message := err.Error()
	if len(msg) > 0 {
		message = msg[0]
	}
	return &ClientError{
		err:     err,
		message: message,
	}
}

func (se *ClientError) Error() string {
	return se.err.Error()
}

func (se *ClientError) Message() string {
	return se.message
}
