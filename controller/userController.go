package controller

import (
	"fmt"
	"net/http"

	"github.com/LIOU2021/gin-layout/db"
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
			"code":    404,
			"message": "data empty",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
	})
}

func (controller *userController) Create(c *gin.Context) {
	user := model.User{UserName: "tester", Password: "12333"}
	conn := db.Conn()
	db.Migrate(conn, &model.User{})

	result := conn.Create(&user)

	if result.Error != nil {
		fmt.Println("Create failt")
		c.JSON(http.StatusBadRequest, "Create failt")
		return
	}

	c.JSON(http.StatusOK, user)
}
