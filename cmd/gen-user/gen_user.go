// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package main

import (
	"fmt"

	"github.com/marcellowy/wow-backup-backend/common"
)

func main() {
	fmt.Println(common.PasswordSalt("aabbccddeeffgghh", "1234567890"))
}
