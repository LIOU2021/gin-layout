package core

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/LIOU2021/gin-layout/env"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var endPoint string

var db *gorm.DB

func initDB() {
	// 組合sql連線字串
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", env.DbSetting.UserName, env.DbSetting.Password, env.DbSetting.Addr, env.DbSetting.Port, env.DbSetting.Database)
	// 連接MySQL
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to mysql failed:", err)
		return
	}
	// 設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	db, err1 := conn.DB()
	if err1 != nil {
		log.Fatal("get db failed:", err)
		return
	}
	db.SetConnMaxLifetime(time.Duration(env.DbSetting.MaxLifetime) * time.Second)
	db.SetMaxIdleConns(env.DbSetting.MaxIdleConns)
	db.SetMaxOpenConns(env.DbSetting.MaxOpenConns)
}

func iniEnv() {
	endPoint = env.AppSetting.HOST + ":" + env.AppSetting.PORT
}

func setTimeZone() {
	os.Setenv("TZ", env.AppSetting.TimeZone)
}

func appInfo() {
	fmt.Println("run " + endPoint)
}

func SetupRouter() *gin.Engine {
	iniEnv()
	setTimeZone()

	gin.SetMode(env.ServerSetting.RunMode)

	router := gin.New()

	register(router)

	return router
}

func ListenAndServe(router *gin.Engine) {
	readTimeout := env.ServerSetting.ReadTimeout
	writeTimeout := env.ServerSetting.WriteTimeout

	s := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	appInfo()
	initDB()
	s.ListenAndServe()
}
