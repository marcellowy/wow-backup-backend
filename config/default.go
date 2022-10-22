// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/20

package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	File  string       // 配置文件
	Viper *viper.Viper // 内容
	Db    *gorm.DB     // database句柄
)
