package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/LIOU2021/gin-layout/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	fileName := "gin.log"
	f := openFile(fileName)

	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())

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
