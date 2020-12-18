package errortype

//RequestValidationError type
type RequestValidationError struct {
	Errors []string `json:"errors"`
}
