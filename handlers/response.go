package handlers

type ResponseMessage string

const (
	success = ResponseMessage("success")
	failed  = ResponseMessage("failed")
)

type Response struct {
	StatusCode int             `json:"code"`
	Message    ResponseMessage `json:"message"`
	Details    interface{}     `json:"details"`
}

func NewResponse(statusCode int, message ResponseMessage, data interface{}) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    message,
		Details:    data,
	}
}
