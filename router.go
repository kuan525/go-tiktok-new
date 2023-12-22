package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"go-tiktok-new/biz/handler"
	"go-tiktok-new/biz/router"
)

func pingRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
}

func register(r *server.Hertz) {
	router.GenerateRegister(r)

	pingRegister(r)
}
