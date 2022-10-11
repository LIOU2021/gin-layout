package main

import (
	"github.com/LIOU2021/gin-layout/config"
	"github.com/LIOU2021/gin-layout/middleware"
	"github.com/LIOU2021/gin-layout/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.Register()
}

func main() {
	router := gin.New()

	middleware.Register(router)
	routes.Register(router)

	router.Run(":8080")
}
