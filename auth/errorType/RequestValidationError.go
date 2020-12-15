package errortype

//RequestValidationError type
type RequestValidationError struct {
	Errors []string `json:"errors"`
}

//ToUniversal transorfs it to universal format
func (err RequestValidationError) ToUniversal(field string) Universal {
	newError := []ErrorModel{}
	for _, a := range err.Errors {
		newError = append(newError, ErrorModel{
			Field:   field,
			Message: a,
		})
	}
	newUniversal := Universal{
		Errors: newError,
	}
	return newUniversal
}
