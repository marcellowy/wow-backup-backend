// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package backup

import (
	"os"

	"github.com/marcellowy/wow-backup-backend/common"
	"github.com/marcellowy/wow-backup-backend/config"
	"github.com/marcellowy/wow-backup-backend/dao"
	"github.com/marcellowy/wow-backup-backend/log"
	"github.com/marcellowy/wow-backup-backend/model"

	"github.com/gin-gonic/gin"
)

type deleteRequest struct {
	BackupID string `json:"backup_id"`
}

// Delete 删除备份
func Delete(ctx *gin.Context) {
	var (
		dr     = deleteRequest{}
		err    error
		backup *model.Backup
	)

	if err = common.ParseBody(ctx, &dr); err != nil {
		common.ResponseErrorJSON(ctx, "读取body出错")
		return
	}

	// 权限校验
	if err = common.Permission(ctx); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseJSON(ctx, common.CodeNoPermission, "无权限", nil)
		return
	}

	if backup, err = dao.NewBackup().QueryByBackupID(config.Db, dr.BackupID); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "查询备份出错")
		return
	}

	if backup.UserID != ctx.GetHeader(common.ConstHeaderUserIDKey) {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "不能删除别人的数据")
		return
	}

	// 删除数据库记录
	if err = dao.NewBackup().DeleteByBackupID(config.Db, dr.BackupID); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "删除数据出错")
		return
	}

	// 检查是不是还有关联到这个hash文件的记录
	// TODO: 有条件可以做成异步处理
	var (
		list       []*model.Backup
		deleteFlag = true // 默认可以删除
	)
	if list, err = dao.NewBackup().QueryListByUserID(config.Db, backup.UserID); err != nil {
		for _, val := range list {
			if val.BackupID != backup.BackupID && val.Hash == backup.Hash {
				// 不能删除
				deleteFlag = false
			}
		}
	}

	// 删除文件
	if deleteFlag {
		if err = os.Remove(backup.RealPath); err != nil {
			// 删除文件不做强制判断,应该需要独立的线程去完成数据库数据与实际文件的对应校验
			log.Warn(ctx, err.Error())
		}
	}
	common.ResponseSuccess(ctx)
}
