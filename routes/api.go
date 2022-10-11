package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// register all api router
func Register(router *gin.Engine) *gin.Engine {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome !")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}
