package errortype

//Universal type
type Universal struct {
	Errors []ErrorModel `json:"errors"`
}

//ErrorModel type
type ErrorModel struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
