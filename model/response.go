package model

import (
	"strconv"

	"github.com/bububa/meituan/util"
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
	// Des 异常描述信息
	Des string `json:"des,omitempty"`
	// Msg 异常描述信息
	Msg string `json:"msg,omitempty"`
	// Status 状态值，0为成功，非0为异常
	Status int `json:"status,omitempty"`
	// Code 状态值，0为成功，非0为异常
	Code int `json:"code,omitempty"`
}

// IsError implement Response interface
func (r BaseResponse) IsError() bool {
	return r.Status != 0 || r.Code != 0
}

// Error implement Response interface
func (r BaseResponse) Error() string {
	if r.Code != 0 {
		return util.StringsJoin(strconv.Itoa(r.Code), ":", r.Msg)
	}
	return util.StringsJoin(strconv.Itoa(r.Status), ":", r.Des)
}
