package main

import (
	"io"
	"os"

	"github.com/LIOU2021/gin-layout/middleware"
	"github.com/LIOU2021/gin-layout/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fileName := "gin.log"
	f := openFile(fileName)

	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()

	middleware.Register(router)
	routes.Register(router)

	router.Run(":8080")
}

func openFile(fileName string) *os.File {
	_, error := os.Stat(fileName)

	if os.IsNotExist(error) {
		f, _ := os.Create(fileName)
		return f
	} else {
		f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		return f
	}
}
