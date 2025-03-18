package model

import (
	"strconv"

	"github.com/bububa/meituan/v2/util"
)

// Response api response interface
type Response interface {
	// IsError 是否返回错误
	IsError() bool
	// Error implement error interface
	Error() string
}

// BaseResponse shared api response data fields
type BaseResponse struct {
	// Code 状态值，0为成功，非0为异常
	Code int `json:"code,omitempty"`
	// Message 响应文案
	Message string `json:"Message,omitempty"`
}

// IsError implement Response interface
func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

// Error implement Response interface
func (r BaseResponse) Error() string {
	return util.StringsJoin(strconv.Itoa(r.Code), ":", r.Message)
}
