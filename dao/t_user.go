// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/20

package dao

import (
	"github.com/marcellowy/wow-backup-backend/model"

	"gorm.io/gorm"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

// Add 添加用户
func (*User) Add(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

// QueryByUserID 查询
func (*User) QueryByUserID(db *gorm.DB, userID string) (*model.User, error) {
	var (
		user = model.User{}
		err  error
	)
	err = db.First(&user, map[string]interface{}{
		"user_id": userID,
	}).Error
	return &user, err
}

// QueryByPhone 查询
func (*User) QueryByPhone(db *gorm.DB, phone string) (*model.User, error) {
	var (
		user = model.User{}
		err  error
	)
	err = db.First(&user, map[string]interface{}{
		"phone": phone,
	}).Error
	return &user, err
}
