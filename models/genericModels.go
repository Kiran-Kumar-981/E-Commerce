package models

const (
	ErrorTypeFatal        = "Fatal"
	ErrorTypeError        = "Error"
	ErrorTypeValidation   = "Validation Error"
	ErrorTypeInfo         = "Info"
	ErrorTypeDebug        = "Debug"
	ErrorTypeUnauthorized = "Unauthorized"
)

type Response struct {
	Error   []ErrorDetails `json:"error,omitempty"`
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data,omitempty"`
}

type ErrorDetails struct {
	ErrorType    string `json:"errorType"`
	ErrorMessage string `json:"errorMessage"`
}
