// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/20

package model

import "time"

// Backup 备份表
type Backup struct {
	ID        uint64    `gorm:"column:id" json:"id"`
	BackupID  string    `gorm:"column:backup_id" json:"backup_id"`
	UserID    string    `gorm:"column:user_id" json:"user_id"`
	Game      string    `gorm:"column:game" json:"game"`
	Name      string    `gorm:"column:name" json:"name"`
	RealPath  string    `gorm:"column:real_path" json:"real_path"`
	Size      uint64    `gorm:"column:size" json:"size"`
	Hash      string    `gorm:"column:hash" json:"hash"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName 表名
func (*Backup) TableName() string {
	return "t_backup"
}
