package goconv

func ErrorToString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func ErrorToStringRef(err error) *string {
	if err == nil {
		return nil
	}
	s := err.Error()
	return &s
}
