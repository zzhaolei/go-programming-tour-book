package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zzhaolei/go-programming-tour-book/blog_service/docs"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/middleware"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/routers/api"
	v1 "github.com/zzhaolei/go-programming-tour-book/blog_service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Translations())
	// OpenAPI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 文件上传服务
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	// 静态资源
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiV1 := r.Group("/api/v1")
	{
		tag := v1.NewTag()
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id", tag.Update)
		apiV1.GET("/tags/:id", tag.Get)
		apiV1.GET("/tags", tag.List)

		article := v1.NewArticle()
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}

	return r
}
