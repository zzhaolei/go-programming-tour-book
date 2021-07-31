package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/service"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/app"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValidate(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errResp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errResp)
		return
	}

	s := service.New(c.Request.Context())
	err := s.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("s.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
