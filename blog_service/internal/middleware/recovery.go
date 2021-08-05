package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/app"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/email"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.New(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err: %v"
				global.Logger.WithCallersFrames().Errorf(s, err)

				err := defaultMailer.Send(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Fatalf("mail.Send err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
