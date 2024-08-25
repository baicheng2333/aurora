package errors

import (
	"errors"
)

type AuroraErrorStatus struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
}

func (a *AuroraErrorStatus) Error() string {
	return a.Message
}

func NewErrors(status uint16, message string) *AuroraErrorStatus {
	return &AuroraErrorStatus{status, message}
}

func WrapStatus(err error) *AuroraErrorStatus {
	var e *AuroraErrorStatus
	if errors.As(err, &e) {
		return e
	}
	return StatusUnWarpError
}

var (
	UsernamePasswordNotMatchStatus = NewErrors(10100, "用户名密码错误!")
)

var (
	StatusUnWarpError     = NewErrors(50500, "未经检查异常信息,请联系管理人员检查日志")
	ParamParseErrorStatus = NewErrors(50501, "参数校验错误,请联系管理人员")
)
