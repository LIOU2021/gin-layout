package db

import (
	"fmt"
	"log"
	"time"

	"net/url"

	"github.com/LIOU2021/gin-layout/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connect *gorm.DB

// get db connection
func Conn() *gorm.DB {
	return connect
}

func setConn(obj *gorm.DB) {
	connect = obj
}

// init db connection
func InitDB() {
	// 組合sql連線字串
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s", env.DbSetting.UserName, env.DbSetting.Password, env.DbSetting.Addr, env.DbSetting.Port, env.DbSetting.Database, url.QueryEscape(env.AppSetting.TimeZone))
	// 連接MySQL
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})

	if err != nil {
		log.Fatal("connection to mysql failed:", err)
		return
	}

	setConn(conn)

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

// exec migrate
func Migrate(obj *gorm.DB, model interface{}) {
	migrator := obj.Migrator()
	has := migrator.HasTable(model)
	if !has {
		fmt.Println("table not exist")
		//創表
		obj.AutoMigrate(model)
	}
}
