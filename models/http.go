package models

type ServiceResponse struct {
	Status       bool        `json:"status"`
	StatusCode   string      `json:"statusCode"`
	ServiceCode  string      `json:"serviceCode"`
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    string      `json:"errorCode,omitempty"`
	ErrorName    string      `json:"errorName,omitempty"`
	ErrorMessage string      `json:"errorMessage,omitempty"`
}
