package config

import (
	"io"

	// "github.com/LIOU2021/gin-layout/core"
	"github.com/LIOU2021/gin-layout/helpers"
	"github.com/gin-gonic/gin"
)

func systemLog() {
	// fileName := "gin.log"
	fileName := LogSetting.FileName
	f := helpers.OpenFile(fileName)

	gin.DefaultWriter = io.MultiWriter(f)
}
