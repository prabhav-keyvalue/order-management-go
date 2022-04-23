package model

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error *ApiError   `json:"error,omitempty"`
}

type ApiError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
