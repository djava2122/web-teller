package service

// response .
type response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Token   string      `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func newResponse(code string, message string) response {
	return response{
		Code:    code,
		Message: message,
	}
}

func getData(data map[string]interface{}, key string) string {
	d, e := data[key]
	if e {
		return d.(string)
	}
	return ""
}
