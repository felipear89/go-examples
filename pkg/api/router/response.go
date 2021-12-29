package router

type ErrorResponse struct {
	Validations []*BodyValidationResponse `json:"validations,omitempty"`
	Error       string                    `json:"error"`
}

type BodyValidationResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}
