package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
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

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", global.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	log.Fatalln(err)
}
