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

type QueryByBackupIDRequest struct {
	BackupID string `json:"backup_id"`
}

// QueryByBackupID 查询列表
func QueryByBackupID(ctx *gin.Context) {

	var (
		err     error
		backup  *model.Backup
		request = QueryByBackupIDRequest{}
	)

	// 权限校验
	if err = common.Permission(ctx); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseJSON(ctx, common.CodeNoPermission, "无权限", nil)
		return
	}

	if err = common.ParseBody(ctx, &request); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "读取body出错")
		return
	}

	if backup, err = dao.NewBackup().QueryByBackupID(config.Db, request.BackupID); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "查询错误")
		return
	}

	if backup.UserID != ctx.GetHeader(common.ConstHeaderUserIDKey) {
		// 这个人查的不是自己的,不输出
		log.Error(ctx, "越权,理论上不应该存在")
		common.ResponseErrorJSON(ctx, "兄弟, 你越权了")
		return
	}

	log.Info(ctx, "success")
	common.ResponseJSON(ctx, 0, "success", backup)
	return
}
