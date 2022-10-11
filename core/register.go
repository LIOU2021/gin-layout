package core

import (
	"github.com/LIOU2021/gin-layout/middleware"
	"github.com/LIOU2021/gin-layout/routes"
	"github.com/gin-gonic/gin"
)

// register gin setting
func Register(router *gin.Engine) *gin.Engine {

	middleware.Register(router)
	routes.Register(router)

	return router
}
