package config

import (
	"io"
	"time"

	"github.com/LIOU2021/gin-layout/env"
	"github.com/LIOU2021/gin-layout/helpers"
	"github.com/gin-gonic/gin"
)

func systemLog() {
	loc, _ := time.LoadLocation(env.AppSetting.TimeZone)
	currentTime := time.Now().In(loc)
	currentDate := currentTime.Format("2006-01-02")
	fileName := "./logs/" + currentDate + "_" + env.LogSetting.FileName
	f := helpers.OpenFile(fileName)

	gin.DefaultWriter = io.MultiWriter(f)
}
