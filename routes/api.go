package routes

import (
	"github.com/LIOU2021/gin-layout/controller"
	"github.com/LIOU2021/gin-layout/middleware"
	"github.com/gin-gonic/gin"
)

// register all api router
func Register(router *gin.Engine) *gin.Engine {
	router.GET("csrf-token", controller.CsrfTokenController.Generate)
	router.GET("verify-csrf-token", middleware.CsrfTokenMiddleware, controller.CsrfTokenController.Verify)

	router.GET("/", controller.IndexController.Index)

	users := router.Group("users")
	users.POST("/", controller.UserController.Create)
	users.GET("/", controller.UserController.Index)

	return router
}
