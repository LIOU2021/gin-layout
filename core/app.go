package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LIOU2021/gin-layout/env"
	"github.com/gin-gonic/gin"
)

var endPoint string

func iniEnv() {
	endPoint = env.AppSetting.HOST + ":" + env.AppSetting.PORT
}

func appInfo() {
	fmt.Println("run " + endPoint)
}

func Run() {
	iniEnv()

	gin.SetMode(env.ServerSetting.RunMode)

	router := gin.New()

	register(router)

	readTimeout := env.ServerSetting.ReadTimeout
	writeTimeout := env.ServerSetting.WriteTimeout

	s := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	appInfo()

	s.ListenAndServe()
}
