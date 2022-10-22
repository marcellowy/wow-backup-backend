// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package common

import (
	"fmt"

	"github.com/marcellowy/wow-backup-backend/config"
	"github.com/marcellowy/wow-backup-backend/dao"
	"github.com/marcellowy/wow-backup-backend/log"

	"github.com/gin-gonic/gin"
)

// permission 判断是否有权限
// 相对简单的一个判断
func permission(ctx *gin.Context) (bool, error) {

	userID := ctx.GetHeader(ConstHeaderUserIDKey)
	password := ctx.GetHeader(ConstHeaderToken)

	if userID == "" || password == "" {
		return false, nil
	}

	user, err := dao.NewUser().QueryByUserID(config.Db, userID)
	if err != nil {
		log.Error(ctx, err.Error())
		return false, err
	}
	if user.Password == password {
		return true, nil
	}
	return false, nil
}

func Permission(ctx *gin.Context) error {
	// 权限校验
	if boo, err := permission(ctx); err != nil {
		log.Error(ctx, "无法判断权限")
		return fmt.Errorf("无法判断权限")
	} else if !boo && err == nil {
		// 无权限
		log.Warn(ctx, "no permission")
		return fmt.Errorf("无权限")
	}
	return nil
}

// PasswordSalt 将密码加盐并hash
func PasswordSalt(p, salt string) (string, error) {
	return Md5(p + salt)
}
