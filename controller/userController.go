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
	//檢索users全部資料
	c.JSON(http.StatusOK, "userController @Index !")
}

func (controller *userController) Create(c *gin.Context) {
	user := model.User{UserName: "tester", Password: "12333"}
	db := db.Conn()

	// db.AutoMigrate(&model.User{})

	// 判斷有沒有table存在
	migrator := db.Migrator()
	has := migrator.HasTable(&model.User{})
	if !has {

		fmt.Println("table not exist")
		//創表
		db.AutoMigrate(&model.User{})
	}

	result := db.Create(&user)

	if result.Error != nil {
		fmt.Println("Create failt")
	}
	// c.JSON(http.StatusOK, user)
	c.JSON(http.StatusOK, "create success !")
}
