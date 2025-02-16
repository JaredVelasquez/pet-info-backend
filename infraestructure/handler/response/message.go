package response

// ErrorResponse representa una estructura para respuestas de error
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse representa una estructura para respuestas exitosas
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
