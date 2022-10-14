package config

import (
	"io"

	"github.com/LIOU2021/gin-layout/env"
	"github.com/LIOU2021/gin-layout/helpers"
	"github.com/gin-gonic/gin"
)

// create log file and set gin.DefaultWriter
func CreateLogFile(fileName string) {
	f := helpers.OpenFile(fileName)
	gin.DefaultWriter = io.MultiWriter(f)
}

// return current log name should be ...
func LogName() string {
	fileName := "./logs/" + env.LogSetting.FileName
	return fileName
}
