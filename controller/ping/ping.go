// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package ping

import (
	"github.com/marcellowy/wow-backup-backend/common"

	"github.com/gin-gonic/gin"
)

// Ping 网络检测接口
func Ping(ctx *gin.Context) {
	common.ResponseJSON(ctx, 0, "pong", nil)
}
