// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/22

package backup

import (
	"net/http"
	"os"

	"github.com/marcellowy/wow-backup-backend/common"
	"github.com/marcellowy/wow-backup-backend/config"
	"github.com/marcellowy/wow-backup-backend/dao"
	"github.com/marcellowy/wow-backup-backend/log"
	"github.com/marcellowy/wow-backup-backend/model"

	"github.com/gin-gonic/gin"
)

type DownloadRequest struct {
	BackupID string `json:"backup_id"`
}

// Download 下载备份文件
func Download(ctx *gin.Context) {
	var (
		ok      bool
		err     error
		backup  *model.Backup
		request = DownloadRequest{}
		byteArr []byte
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

	// 判断文件是不是属于这个人
	if backup.UserID != ctx.GetHeader(common.ConstHeaderUserIDKey) {
		// 这个人查的不是自己的,不输出
		log.Error(ctx, "越权,理论上不应该存在")
		common.ResponseErrorJSON(ctx, "兄弟, 你越权了")
		return
	}

	if ok, err = common.PathExists(backup.RealPath); err != nil || !ok {
		// 文件有问题
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "文件存在问题, 无法下载")
		return
	}

	if byteArr, err = os.ReadFile(backup.RealPath); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "无法读取文件")
		return
	}

	// 开始下载文件
	fileContentDisposition := "attachment;filename=\"" + backup.Name + "\""
	ctx.Header("Content-Type", "application/x-7z-compressed") // 这里是压缩文件类型 .zip
	ctx.Header("Content-Disposition", fileContentDisposition)
	ctx.Data(http.StatusOK, "application/x-7z-compressed", byteArr)
}
