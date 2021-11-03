package service

// response .
type response2 struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Token   string      `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func newResponse(code string, message string) response2 {
	return response2{
		Code:    code,
		Message: message,
	}
}

func newResponseWithData(code string, message string, data interface{}) response2 {
	return response2{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
func respCekStatus(data interface{}, code, message string) response2 {
	return response2{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
func successResp(data interface{}) response2 {
	return response2{
		Code:    "00",
		Message: "Success",
		Data:    data,
	}
}

func getData(data map[string]interface{}, key string) string {
	d, e := data[key]
	if e {
		return d.(string)
	}
	return ""
}

type ResSiskopatuh struct {
	RC      string `json:"RC"`
	Message string `json:"message"`
	Token   string `json:"token"`
}