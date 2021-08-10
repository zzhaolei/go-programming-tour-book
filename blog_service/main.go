package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/tracer"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/model"

	"github.com/gin-gonic/gin"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/setting"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/routers"
)

var (
	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
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

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}

	err = setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
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

	err = s.ReadSection("Email", &global.EmailSetting)
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

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTrace(
		"blog_service",
		"127.0.0.1:6831",
	)

	if err != nil {
		return err
	}

	global.Tracer = jaegerTracer
	return nil
}

func setupFlag() error {
	flag.BoolVar(&isVersion, "version", false, "编译信息")
	flag.Parse()
	return nil
}

// main 入口函数
// @title 博客系统
// @version 1.0.0
// @description Go语言编程之旅
// @termsOfService 没有呢
func main() {
	if isVersion {
		fmt.Println("BuildTime: ", buildTime)
		fmt.Println("BuildVersion: ", buildVersion)
		fmt.Println("GitCommitID: ", gitCommitID)
		return
	}
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

	go func() {
		log.Println("ListenAndServe: ", addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("server.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server Exit.")
}
