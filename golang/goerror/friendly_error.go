package goerror

import "fmt"

// FriendlyError реализация интерфейса interfaces.FriendlyError
type FriendlyError struct {
	error   error
	message string
}

var _ error = (*FriendlyError)(nil)

func NewFriendlyError(err error, message string) *FriendlyError {
	return &FriendlyError{
		error:   err,
		message: message,
	}
}

func MakeFriendlyErrorf(friendly, message string, a ...any) *FriendlyError {
	return &FriendlyError{
		fmt.Errorf(message, a...),
		friendly,
	}
}

func (e FriendlyError) Error() string {
	return e.error.Error()
}

func (e FriendlyError) Friendly() string {
	return e.message
}
