package main

import (
	"log"
	"net/http"
	"time"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/internal/routers"
)

func main() {
	router := routers.NewRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	log.Fatalln(err)
}
