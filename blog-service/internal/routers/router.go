package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zzhaolei/go-programming-tour-book/blog_service/docs"
	v1 "github.com/zzhaolei/go-programming-tour-book/blog_service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 文档
	url := ginSwagger.URL("http://127.0.0.1:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//
	tag := v1.NewTag()
	article := v1.NewArticle()
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.GET("/tags", tag.List)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.GET("/tags/:id", tag.Get)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PATCH("/tags/:id/state", tag.Update)

		apiV1.POST("/articles", article.Create)
		apiV1.GET("/articles", article.List)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PATCH("/articles/:id/state", article.Update)
	}
	return r
}
