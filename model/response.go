package model

type Response struct {
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Error    *ApiError   `json:"error,omitempty"`
	Errors   []ApiError  `json:"errors,omitempty"`
	PageInfo *PageInfo   `json:"pageInfo,omitempty"`
}

type ApiError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
