package response

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SUCCESS
func Success(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

// ERROR
func Error(msg string) Response {
	return Response{
		Success: false,
		Error:   msg,
	}
}