package core

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/LIOU2021/gin-layout/db"
	"github.com/LIOU2021/gin-layout/env"
	"github.com/gin-gonic/gin"
)

var endPoint string

func iniEnv() {
	endPoint = env.AppSetting.HOST + ":" + env.AppSetting.PORT
}

func setTimeZone() {
	os.Setenv("TZ", env.AppSetting.TimeZone)
}

func appInfo() {
	fmt.Println("run " + endPoint)
}

func SetupRouter() *gin.Engine {
	iniEnv()
	setTimeZone()

	gin.SetMode(env.ServerSetting.RunMode)

	router := gin.New()

	register(router)

	return router
}

func ListenAndServe(router *gin.Engine) {
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
	db.InitDB()
	s.ListenAndServe()
}
