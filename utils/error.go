package utils

type ServerError struct {
	Name    string
	Message string
}

func (e *ServerError) Error() string {
	return e.Message
}
