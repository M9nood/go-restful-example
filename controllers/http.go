package controllers

import (
	"m9hub/errors"
	"m9hub/models"
)

var serviceCode = "00"

func CreateSuccessResponse(data interface{}) models.ServiceResponse {
	return models.ServiceResponse{
		Status:      true,
		StatusCode:  "200",
		ServiceCode: serviceCode,
		Data:        data,
	}
}

// CreateErrorResponse - create fail response
func CreateErrorResponse(data interface{}, errName string, msg string) models.ServiceResponse {
	errorCode, errorName := errors.ErrorHandler(errName)
	return models.ServiceResponse{
		Status:       false,
		ServiceCode:  serviceCode,
		Data:         data,
		ErrorCode:    errorCode,
		ErrorName:    errorName,
		ErrorMessage: msg,
	}
}
