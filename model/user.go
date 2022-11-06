package model

import (
	"time"
)

type User struct {
	ID        int64      `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	UserName  string     `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  string     `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	CreatedAt *time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}
