package middleware

import (
	"github.com/LIOU2021/gin-layout/env"
	"github.com/LIOU2021/gin-layout/helpers"
	"github.com/gin-gonic/gin"
)

func CsrfTokenMiddleware(ctx *gin.Context) {
	key := env.AppSetting.Key
	hash := ctx.GetHeader("Csrf-Token")
	check := helpers.HashCheck(hash, key)
	if !check {
		ctx.Abort()
		ctx.JSON(403, gin.H{"msg": "Invalid Csrf-Token"})
	}
}
