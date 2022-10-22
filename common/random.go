// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package common

import (
	"math/rand"
	"time"
)

var ra *rand.Rand

// Contain 随机字符串
type type_ int

const (
	RandomLowercase type_ = 1 // 小写字母
	RandomMajuscule       = 2 // 大写字母
	RandomNumber          = 4 // 数字
	RandomSymbol          = 8 // 符号
	RandomAll             = RandomLowercase | RandomMajuscule | RandomNumber | RandomSymbol
)

var types = []type_{RandomLowercase, RandomMajuscule, RandomNumber, RandomSymbol}

var typesValue = map[type_][]byte{
	RandomLowercase: lowercase,
	RandomMajuscule: majuscule,
	RandomNumber:    number,
	RandomSymbol:    symbol,
}

// lowercase 小写字母
var lowercase = []byte{
	0x61, 0x62, 0x63, 0x64, 0x65, 0x66,
	0x67, 0x68, 0x69, 0x6A, 0x6B, 0x6D,
	0x6E, 0x70, 0x71, 0x72, 0x73, 0x74,
	0x75, 0x76, 0x77, 0x78, 0x79, 0x7A,
}

// majuscule 大写字母
var majuscule = []byte{
	0x41, 0x42, 0x43, 0x44, 0x45, 0x46,
	0x47, 0x48, 0x49, 0x4A, 0x4B, 0x4D,
	0x4E, 0x50, 0x51, 0x52, 0x53, 0x54,
	0x55, 0x56, 0x57, 0x58, 0x59, 0x5A,
}

// number 数字
var number = []byte{'2', '3', '4', '5', '6', '7', '8', '9'}

// symbol 符号
var symbol = []byte{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '-', '+'}

// init 初始化
func init() {
	ra = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomString 随机字符串
func RandomString(length uint, contain type_) string {

	var (
		result string
		r, t   []byte
		i      uint
	)

	for _, v := range types {
		if v&contain > 0 {
			r = append(r, typesValue[v]...)
		}
	}

	if len(r) == 0 {
		return ""
	}

	var l = len(r)
	for i = 0; i < length; i++ {
		t = append(t, r[ra.Intn(l-1)])
	}
	result = string(t)

	return result
}
