package main

import (
	"fmt"

	"github.com/LIOU2021/gin-layout/config"
	"github.com/LIOU2021/gin-layout/core"
	"github.com/gin-gonic/gin"
)

func init() {
	core.Env()
	config.Kernel()
}

func main() {
	router := gin.New()
	core.Register(router)
	fmt.Println("run http://127.0.0.1:8080")
	router.Run("127.0.0.1:8080")
}
