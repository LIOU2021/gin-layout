package main

import (
	"fmt"

	"github.com/LIOU2021/gin-layout/config"
	"github.com/LIOU2021/gin-layout/core"
	"github.com/LIOU2021/gin-layout/env"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Kernel()
}

func main() {
	router := gin.New()
	core.Register(router)

	endPoint := env.AppSetting.HOST + ":" + env.AppSetting.PORT
	fmt.Println("run " + endPoint)
	router.Run(endPoint)
}
