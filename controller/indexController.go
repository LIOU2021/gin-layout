package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type indexController struct{}

var IndexController = &indexController{}

func (controller *indexController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, "welcome !")
}
