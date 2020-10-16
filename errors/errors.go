package errors

import "fmt"

type HttpError struct {
	StatusCode  int
	HttpMethod  string
	ResourceUrl string
	Message     string
}

func New(message string, statusCode int, httpMethod string, url string) *HttpError {
	return &HttpError{
		StatusCode:  statusCode,
		HttpMethod:  httpMethod,
		ResourceUrl: url,
		Message:     message,
	}
}

func FromError(err error, statusCode int, httpMethod string, url string) *HttpError {
	return &HttpError{
		StatusCode:  statusCode,
		HttpMethod:  httpMethod,
		ResourceUrl: url,
		Message:     err.Error(),
	}
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%s request to %s failed with %d. Error %s",
		e.HttpMethod, e.ResourceUrl, e.StatusCode, e.Message)
}
