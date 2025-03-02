package transport

type APIErrorResponse struct {
	Description      string   `json:"description"`
	Code             string   `json:"code"`
	ExceptionName    string   `json:"exception_name"`
	ExceptionMessage string   `json:"exception_message"`
	StackTrace       []string `json:"stack_trace"`
}
