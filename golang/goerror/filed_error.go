package goerror

type ErrorFiled struct {
	error
	message   string
	fieldName string
}

func NewFiledError(err error, message string, fieldName string) ErrorFiled {
	return ErrorFiled{
		error:     err,
		message:   message,
		fieldName: fieldName,
	}
}

func (e ErrorFiled) Friendly() string {
	return e.message
}

func (e ErrorFiled) Field() string {
	return e.fieldName
}
