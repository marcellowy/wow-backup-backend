// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package common

import (
	"fmt"
	"os"
)

// PathExists 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	// isnotexist来判断，是不是不存在的错误
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		return false, nil
	}
	return false, err //如果有错误了，但是不是不存在的错误，所以把这个错误原封不动的返回
}

// CreateDirectory 创建目录
func CreateDirectory(path string) error {

	var (
		ok  bool
		err error
	)

	if ok, err = PathExists(path); ok && err == nil {
		// 目录存在
		return nil
	} else if !ok && err == nil {
		// 文件不存在
		if err = os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("创建目录失败" + err.Error())
		}
		return nil
	} else if !ok && err != nil {
		// 判断时出现了系统错误
		return err
	}
	return nil
}
