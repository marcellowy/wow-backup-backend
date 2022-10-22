// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func FileMd5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Md5(str string) (string, error) {
	w := md5.New()
	if _, err := io.WriteString(w, str); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", w.Sum(nil)), nil //w.Sum(nil)将w的hash转成[]byte格式
}
