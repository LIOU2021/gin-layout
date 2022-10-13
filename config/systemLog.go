package config

import (
	"io"
	"time"

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
	loc, _ := time.LoadLocation(env.AppSetting.TimeZone)
	currentTime := time.Now().In(loc)
	currentDate := currentTime.Format("2006-01-02")
	fileName := "./logs/" + currentDate + "_" + env.LogSetting.FileName
	return fileName
}
