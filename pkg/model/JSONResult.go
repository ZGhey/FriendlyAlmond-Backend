package model

import (
	"FriendlyAlmond_backend/pkg/utils"
	"fmt"
)

type JSONResult struct {
	Code    string `json:"code" example:"0"`
	Message string `json:"message" example:"success"`
	Success bool   `json:"success" example:"true"`
}

func (j *JSONResult) GetMessage() string {
	return j.Message
}
func (j *JSONResult) NewError(recode string) *JSONResult {
	j.Code = recode
	j.Message = utils.RecodeTest(recode)
	j.Success = false
	return j
}

func (j *JSONResult) NewSuccess() *JSONResult {
	j.Code = utils.RECODE_OK
	j.Message = utils.RecodeTest(utils.RECODE_OK)
	j.Success = true
	return j
}

func (j *JSONResult) SetError(errorCode string, message string, err error) string {
	j.Code = errorCode
	j.Message = message
	j.Success = false
	j.Message = fmt.Sprintf("message is: %s - %s, extra error message: %v",
		utils.RecodeTest(errorCode), message, err)
	return j.Message
}
