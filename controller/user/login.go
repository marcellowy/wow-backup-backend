// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package user

import (
	"github.com/marcellowy/wow-backup-backend/common"
	"github.com/marcellowy/wow-backup-backend/config"
	"github.com/marcellowy/wow-backup-backend/dao"
	"github.com/marcellowy/wow-backup-backend/log"
	"github.com/marcellowy/wow-backup-backend/model"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Login 登录
func Login(ctx *gin.Context) {
	var (
		lr   = loginRequest{}
		user *model.User
		err  error
	)
	if err = common.ParseBody(ctx, &lr); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "读取body出错")
		return
	}
	if user, err = dao.NewUser().QueryByPhone(config.Db, lr.Phone); err != nil {
		log.Warn(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "用户名或密码错误")
		return
	}
	var sp string
	if sp, err = common.PasswordSalt(lr.Password, user.Salt); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "password salt错误")
		return
	}

	if sp != user.Password {
		log.Error(ctx, "password error")
		common.ResponseErrorJSON(ctx, "密码错误")
		return
	}

	log.Info(ctx, "success")
	common.ResponseJSON(ctx, 0, "success", map[string]interface{}{
		"user_id":  user.UserID,
		"token":    user.Password,
		"nickname": user.Nickname,
	})
}
