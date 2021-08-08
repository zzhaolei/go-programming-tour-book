package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zzhaolei/go-programming-tour-book/blog_service/docs"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/middleware"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/routers/api"
	v1 "github.com/zzhaolei/go-programming-tour-book/blog_service/internal/routers/api/v1"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/limit"
)

var methodLimiters = limit.NewMethodLimiter().AddBuckets(limit.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger(), gin.Recovery())
	} else {
		r.Use(middleware.AccessLog(), middleware.Recovery())
	}
	r.Use(
		middleware.AccessLog(),
		middleware.RateLimiter(methodLimiters),
		middleware.ContextTimeout(global.ServerSetting.DefaultContextTimeout*time.Second),
		middleware.Translations(),
		middleware.Tracing(),
	)

	// 登陆
	r.POST("/auth", api.GetAuth)

	// OpenAPI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 文件上传服务
	upload := api.NewUpload()
	r.POST("/upload/file", middleware.JWT(), upload.UploadFile)

	// 静态资源
	static := r.Group("/static", middleware.JWT())
	static.StaticFS("/", gin.Dir(global.AppSetting.UploadSavePath, false))

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
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
