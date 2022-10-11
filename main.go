package main

import (
	"github.com/LIOU2021/gin-layout/config"
	"github.com/LIOU2021/gin-layout/core"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Register()
}

func main() {
	router := gin.New()
	core.Register(router)
	router.Run(":8080")
}
