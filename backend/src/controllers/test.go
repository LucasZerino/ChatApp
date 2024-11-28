package controllers

import (
	"backend/src/configs/rest_error"
)

func Test(message string) (err *rest_error.RestErr) {
	return rest_error.NewBadRequestError(message)
}
