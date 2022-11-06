package model

import (
	"time"
)

type User struct {
	//gorm為model的tag標籤，v2版的auto_increment要放在type裡面，v1版是放獨立定義
	ID         int64     `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	UserName   string    `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password   string    `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	Created_at time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	Update_at  time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}
