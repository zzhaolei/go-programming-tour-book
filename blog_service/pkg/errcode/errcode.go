package errcode

import (
	"fmt"
	"net/http"
)

var codes = map[int]string{}

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

func NewError(code int, msg string) *Error {
	if value, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，异常信息为: %s，请更换一个", code, value))
	}
	codes[code] = msg
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Msgf(args ...interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = details
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code:
		fallthrough
	case UnauthorizedTokenError.Code:
		fallthrough
	case UnauthorizedTokenGenerate.Code:
		fallthrough
	case UnauthorizedTokenTimeout.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
