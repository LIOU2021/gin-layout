package routes

import (
	"net/http"

	"github.com/LIOU2021/gin-layout/env"
	"github.com/LIOU2021/gin-layout/helpers"
	"github.com/LIOU2021/gin-layout/middleware"
	"github.com/gin-gonic/gin"
)

// register all api router
func Register(router *gin.Engine) *gin.Engine {
	router.GET("csrf-token", func(ctx *gin.Context) {
		key := env.AppSetting.Key
		hash := helpers.HashMake(key)
		ctx.Header("Csrf-Token", hash)
		ctx.String(http.StatusOK, "success")
	})

	router.GET("verify-csrf-token", middleware.CsrfTokenMiddleware, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "csrf token verify success")
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome !")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ping~~"})
	})

	return router
}
