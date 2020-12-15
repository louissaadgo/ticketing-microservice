package errortype

//DatabaseConnectionError type
type DatabaseConnectionError struct {
	Error string `json:"error"`
}

//ToUniversal transforms it to universal format
func (err DatabaseConnectionError) ToUniversal(field string) Universal {
	newError := []ErrorModel{}
	newError = append(newError, ErrorModel{
		Field:   field,
		Message: err.Error,
	})
	newUniversal := Universal{
		Errors: newError,
	}
	return newUniversal
}
