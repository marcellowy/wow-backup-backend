// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package router

import (
	"github.com/marcellowy/wow-backup-backend/controller/backup"
	"github.com/marcellowy/wow-backup-backend/controller/ping"
	"github.com/marcellowy/wow-backup-backend/controller/user"

	"github.com/gin-gonic/gin"
)

// Router 路由
func Router(engine *gin.Engine) {
	engine.GET("/api/ping", ping.Ping)
	// 备份接口
	engine.POST("/api/backup/list", backup.List)
	engine.POST("/api/backup/upload", backup.Upload)
	engine.POST("/api/backup/delete", backup.Delete)
	engine.POST("/api/backup/download", backup.Download)
	engine.POST("/api/backup/query-by-backup-id", backup.QueryByBackupID)

	// 用户接口
	engine.POST("/api/user/login", user.Login)
}
