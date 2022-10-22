// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/marcellowy/wow-backup-backend/log"

	"github.com/gin-gonic/gin"
)

// ResponseJSON 接口统一返回
func ResponseJSON(ctx *gin.Context, code uint, message string, data interface{}) {
	ctx.JSON(200, map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

// ResponseSuccess 返回成功
func ResponseSuccess(ctx *gin.Context) {
	ctx.JSON(200, map[string]interface{}{
		"code":    0,
		"message": "success",
	})
}

// ResponseErrorJSON 统一错误返回
func ResponseErrorJSON(ctx *gin.Context, message string) {
	ctx.JSON(200, map[string]interface{}{
		"code":    255,
		"message": message,
	})
}

// ParseBody 解析body
func ParseBody(ctx *gin.Context, a interface{}) error {
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Error(ctx, err.Error())
		return fmt.Errorf("读取内容出错")
	}

	if err = json.Unmarshal(data, &a); err != nil {
		log.Error(ctx, err.Error())
		return fmt.Errorf("解析body出错")
	}
	return nil
}
