package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/model"

	"github.com/gin-gonic/gin"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/setting"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/routers"
)

func init() {
	var err error
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func setupSetting() error {
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	var builder strings.Builder
	builder.WriteString(global.AppSetting.LogSavePath)
	builder.WriteString("/")
	builder.WriteString(global.AppSetting.LogFileName)
	builder.WriteString(global.AppSetting.LogFileExt)
	filename := builder.String()
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  filename,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

// main 入口函数
// @title 博客系统
// @version 1.0.0
// @description Go语言编程之旅
// @termsOfService 没有呢
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	addr := fmt.Sprintf(":%s", global.ServerSetting.HttpPort)
	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("ListenAndServe: %s", addr)
	err := server.ListenAndServe()
	log.Fatalln(err)
}
