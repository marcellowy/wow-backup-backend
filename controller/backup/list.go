// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package backup

import (
	"github.com/marcellowy/wow-backup-backend/common"
	"github.com/marcellowy/wow-backup-backend/config"
	"github.com/marcellowy/wow-backup-backend/dao"
	"github.com/marcellowy/wow-backup-backend/log"
	"github.com/marcellowy/wow-backup-backend/model"

	"github.com/gin-gonic/gin"
)

// List 查询列表
func List(ctx *gin.Context) {

	var (
		err    error
		backup []*model.Backup
	)

	// 权限校验
	if err = common.Permission(ctx); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseJSON(ctx, common.CodeNoPermission, "无权限", nil)
		return
	}
	if backup, err = dao.NewBackup().QueryListByUserID(config.Db, ctx.GetHeader(common.ConstHeaderUserIDKey)); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "查询错误")
		return
	}

	log.Info(ctx, "success")
	common.ResponseJSON(ctx, 0, "success", backup)
	return
}
