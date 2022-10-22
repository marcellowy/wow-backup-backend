// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/20

package model

import (
	"time"

	"gorm.io/gorm"
)

// User 备份表
type User struct {
	gorm.Model
	ID            uint64    `gorm:"column:id"`
	UserID        string    `gorm:"column:user_id"`
	Phone         string    `gorm:"column:phone"`
	Password      string    `gorm:"column:password"`
	Salt          string    `gorm:"column:salt"`
	Nickname      string    `gorm:"column:nickname"`
	RegisterTime  time.Time `gorm:"column:register_time"`
	LastLoginTime time.Time `gorm:"column:last_login_time"`
}

// TableName 表名
func (*User) TableName() string {
	return "t_user"
}
