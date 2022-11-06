package model

import (
	"time"

	"github.com/LIOU2021/gin-layout/db"
)

type User struct {
	ID        int64      `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	UserName  string     `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  string     `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	CreatedAt *time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

// var Users []User

// func init() {
// 	test := User{UserName: "hello", Password: "123456"}
// 	db.Migrate(db.Conn(), &test)
// }

// 列表
func (user *User) Users() (users []User, err error) {
	if err = db.Conn().Find(&users).Error; err != nil {
		return
	}
	return
}

// 新增
func (user User) Insert() (id int64, err error) {

	result := db.Conn().Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
