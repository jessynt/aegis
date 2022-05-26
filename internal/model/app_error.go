package model

type AppCoder interface {
	error

	AppMessage() string
	AppCode() int
}

type AppError struct {
	code    int
	message string
}

func MakeAppError(code int, message string) AppError {
	return AppError{code, message}
}

func (a AppError) Error() string {
	return a.message
}

func (a AppError) AppCode() int {
	return a.code
}

func (a AppError) AppMessage() string {
	return a.message
}
