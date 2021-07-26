package errors

type InvalidParamaetersError struct {
	Err string
}

func (m *InvalidParamaetersError) Error() string {
	return m.Err
}
