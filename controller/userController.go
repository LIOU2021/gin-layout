package controller

import (
	"net/http"

	"github.com/LIOU2021/gin-layout/model"
	"github.com/gin-gonic/gin"
)

type userController struct{}

var UserController = &userController{}

func (controller *userController) Index(c *gin.Context) {
	var user model.User
	result, err := user.Users()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "data empty",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func (controller *userController) Create(c *gin.Context) {

	user := model.User{UserName: "hello", Password: "123456"}
	_, err := user.Insert()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "create failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "create success",
		"data":    user,
	})

}
