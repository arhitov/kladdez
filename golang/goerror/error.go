package goerror

import "fmt"

func MakeMapErrorsString(errs []error) map[string]string {
	if errs == nil {
		return nil
	}
	mapErrorsString := map[string]string{}
	for idx, err := range errs {
		var message string
		if e, ok := err.(interface{ Friendly() string }); ok {
			message = e.Friendly()
		} else {
			message = err.Error()
		}
		if e, ok := err.(interface{ Field() string }); ok {
			mapErrorsString[e.Field()] = message
		} else {
			mapErrorsString[fmt.Sprintf("error-%d", idx)] = message
		}
	}
	return mapErrorsString
}

func MakeErrorsString(errs []error) string {
	if errs == nil {
		return ""
	}
	errorsString := ""
	for _, err := range errs {
		var message string
		if e, ok := err.(interface{ Friendly() string }); ok {
			message = e.Friendly()
		} else {
			message = err.Error()
		}
		if message == "" {
			errorsString += "; "
		}
		if e, ok := err.(interface{ Field() string }); ok {
			errorsString += fmt.Sprintf("filed [%s] %s", e.Field(), message)
		} else {
			errorsString += message
		}
	}
	return errorsString
}
