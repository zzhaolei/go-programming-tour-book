package main

import (
	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/model"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/logger"
	settings "github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/settings"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"
	routers "github.com/zzhaolei/go-programming-tour-book/blog_service/internal/routers"
)


func setupSetting() error {
	setting, err := settings.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	// 使用 全局DBEngine 防止包外调用 DBEngine 时此字段被重新创建为 nil
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}


func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.global.SetupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.global.SetupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.global.SetupLogger err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description 学习使用Go语言开发Web后端系统
// @termsOfService 无
func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()

	addr := ":" + global.ServerSetting.HttpPort
	log.Printf("ListenAndServe: %s\n", addr)
	s := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println(s.ListenAndServe())
}
