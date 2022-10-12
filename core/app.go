package core

import (
	"fmt"

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

	router := gin.New()
	register(router)

	appInfo()

	router.Run(endPoint)
}
