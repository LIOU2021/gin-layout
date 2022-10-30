package main

import (
	"github.com/LIOU2021/gin-layout/config"
	"github.com/LIOU2021/gin-layout/core"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Kernel()
}

func main() {
	router := GetRouter()
	core.ListenAndServe(router)
}

func GetRouter() *gin.Engine {
	router := core.SetupRouter()
	return router
}
