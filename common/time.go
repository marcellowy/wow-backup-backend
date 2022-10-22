// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package common

import "time"

// NowDatetime 当前时间
func NowDatetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
