package app

import (
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
)

// ValidateError 验证失败的异常结构体
type ValidateError struct {
	Key     string
	Message string
}

// ValidateErrors 验证失败的异常结构体切片
type ValidateErrors []*ValidateError

// Error 获取异常信息
func (v *ValidateError) Error() string {
	return v.Message
}

// Error 处理多个异常信息
func (v ValidateErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

// Errors 获取异常信息切片
func (v ValidateErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

// BindAndValidate 验证数据
func BindAndValidate(c *gin.Context, v interface{}) (bool, ValidateErrors) {
	var errs ValidateErrors

	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		_errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}

		for key, value := range _errs.Translate(trans) {
			errs = append(errs, &ValidateError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}
