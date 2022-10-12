package config

import (
	"io"

	"github.com/LIOU2021/gin-layout/config/env"
	"github.com/LIOU2021/gin-layout/helpers"
	"github.com/gin-gonic/gin"
)

func systemLog() {
	// fileName := "gin.log"
	fileName := env.LogSetting.FileName
	f := helpers.OpenFile(fileName)

	gin.DefaultWriter = io.MultiWriter(f)
}
