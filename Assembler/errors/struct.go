package errors

type CustomError struct {
	s       string
	errType int
}

func NewCustomError(e error, errType int) *CustomError {
	return &CustomError{s: e.Error(), errType: errType}
}

func (e *CustomError) Error() string {
	return e.s
}

func (e *CustomError) IsCodeError() bool {
	return e.errType == CodeError
}

const AssemblerError = 500
const CodeError = 400
