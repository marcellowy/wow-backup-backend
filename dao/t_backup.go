// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/20

package dao

import (
	"github.com/marcellowy/wow-backup-backend/model"

	"gorm.io/gorm"
)

type Backup struct{}

// NewBackup 实例化
func NewBackup() *Backup {
	return &Backup{}
}

// Add 添加备份记录
func (*Backup) Add(db *gorm.DB, backup *model.Backup) error {
	return db.Create(backup).Error
}

// QueryListByUserID 查询列表
func (*Backup) QueryListByUserID(db *gorm.DB, userID string) ([]*model.Backup, error) {
	var data []*model.Backup
	var err = db.Find(&data, map[string]interface{}{
		"user_id": userID,
	}).Order("created_at DESC").Error
	return data, err
}

// QueryByBackupID 查询
func (*Backup) QueryByBackupID(db *gorm.DB, backupID string) (*model.Backup, error) {
	var data = model.Backup{}
	var err = db.Find(&data, map[string]interface{}{
		"backup_id": backupID,
	}).Error
	return &data, err
}

// DeleteByBackupID 查询
func (*Backup) DeleteByBackupID(db *gorm.DB, backupID string) error {
	return db.Delete(&model.Backup{}, map[string]interface{}{
		"backup_id": backupID,
	}).Error
}
