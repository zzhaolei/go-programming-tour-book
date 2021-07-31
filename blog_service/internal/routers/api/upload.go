package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/service"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/app"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/convert"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/errcode"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	// 读取入参file字段的上传文件信息
	file, fileHeader, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()

	if err != nil {
		errResp := errcode.InvalidParams.WithDetails(err.Error())
		response.ToErrorResponse(errResp)
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	fileInfo, err := service.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf("service.UploadFile err: %v", err)
		errResp := errcode.ErrorUploadFileFail.WithDetails(err.Error())
		response.ToErrorResponse(errResp)
		return
	}

	response.ToResponse(gin.H{
		"file": fileInfo.AccessURL,
	})

}
