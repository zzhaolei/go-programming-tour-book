package app

import (
	"net/http"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/errcode"

	"github.com/gin-gonic/gin"
)

// Response 响应
type Response struct {
	Ctx *gin.Context
}

// Pager 页
type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

// NewResponse 创建Response对象
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// ToResponse 响应
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code":    err.Code,
		"msg":     err.Msg,
		"details": []string{},
	}
	if len(err.Details) > 0 {
		response["details"] = err.Details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
