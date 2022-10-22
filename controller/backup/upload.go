// Copyright (c) 2022 Tencent.Ltd. All rights reserved.
// Author: chadwang@tencent.com
// Date: 2022/10/21

package backup

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/marcellowy/wow-backup-backend/common"
	"github.com/marcellowy/wow-backup-backend/config"
	"github.com/marcellowy/wow-backup-backend/dao"
	"github.com/marcellowy/wow-backup-backend/log"
	"github.com/marcellowy/wow-backup-backend/model"

	"github.com/gin-gonic/gin"
)

// mergePart 合并分片
func mergePart(ctx *gin.Context, tmpDirectory, hash string, total int) (string, string, error) {

	var (
		directory = config.Viper.GetString("backup.save_directory")
		file      = fmt.Sprintf("%s/%s.7z", directory, hash)
		b         []byte
		newHash   string
	)

	fd, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Error(ctx, err.Error())
		return "", newHash, fmt.Errorf("创建文件失败")
	}
	defer fd.Close()

	for i := 1; i <= total; i++ {
		p := getPartPath(tmpDirectory, hash, total, i)
		if b, err = os.ReadFile(p); err != nil {
			log.Error(ctx, err.Error())
			return "", newHash, fmt.Errorf("读取分片错误")
		}
		if _, err = fd.Write(b); err != nil {
			log.Error(ctx, err.Error())
			return "", newHash, fmt.Errorf("写入分片错误")
		}
	}

	// hash校验
	if newHash, err = common.FileMd5(file); err != nil {
		log.Error(ctx, "hash: "+hash+" new Hash: "+newHash+" 不相同")
		return "", newHash, fmt.Errorf("创建hash失败")
	} else if newHash != hash {
		log.Error(ctx, "hash: "+hash+" new Hash: "+newHash+" 不相同")
		return "", newHash, fmt.Errorf("hash校验失败")
	}

	// 删除所有分片
	removePart(ctx, tmpDirectory, hash, total)

	return file, newHash, nil
}

// removePart 删除所有分片
func removePart(ctx *gin.Context, tmpDirectory, hash string, total int) {
	for i := 1; i <= total; i++ {
		p := getPartPath(tmpDirectory, hash, total, i)
		if err := os.Remove(p); err != nil {
			log.Warn(ctx, err.Error())
		}
	}
}

// Upload 上传备份
func Upload(ctx *gin.Context) {

	var (
		tmpDirectory = config.Viper.GetString("backup.tmp_directory") // 临时保存地址
		err          error
		boo          bool
		file7z       string
	)

	// 创建临时目录
	if boo, err = common.PathExists(tmpDirectory); err == nil && !boo {
		if err = os.MkdirAll(tmpDirectory, 0755); err != nil {
			// 创建目录错误
			log.Error(ctx, err.Error())
			common.ResponseErrorJSON(ctx, "创建目录失败")
			return
		}
	} else if err != nil && !boo {
		// 判断时发生了未知错误
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "创建目录发生未知错误")
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "无法接收到文件")
		return
	}

	hash := ctx.PostForm("hash")   // 完整文件的hash
	total := ctx.PostForm("total") // 分片总数
	index := ctx.PostForm("index") // 当前传输分片数
	size := ctx.PostForm("size")   // 完整文件大小
	game := ctx.PostForm("game")   // 游戏
	userID := ctx.GetHeader(common.ConstHeaderUserIDKey)

	// 权限校验
	if err = common.Permission(ctx); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseJSON(ctx, common.CodeNoPermission, "无权限", nil)
		return
	}

	// 将total转为整型
	totalInt, err := strconv.ParseInt(total, 10, 64)
	if err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "total类型不正确")
		return
	}

	if totalInt <= 0 {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "total范围不正确")
		return
	}

	// 将index转为整型
	indexInt, err := strconv.ParseInt(index, 10, 64)
	if err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "index类型不正确")
		return
	}

	if indexInt <= 0 {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "index范围不正确")
		return
	}

	// 将size转为整型
	sizeInt, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "size类型不正确")
		return
	}

	if sizeInt <= 0 {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "size范围不正确")
		return
	}

	// 分片保存路径
	partName := getPartPath(tmpDirectory, hash, int(totalInt), int(indexInt))

	// 保存分片
	if err = ctx.SaveUploadedFile(file, partName); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "保存分片文件出错")
		return
	}

	if totalInt != indexInt {
		log.Info(ctx, "success")
		common.ResponseJSON(ctx, 0, "success", nil)
		return
	}

	// 最后一个分片
	if file7z, hash, err = mergePart(ctx, tmpDirectory, hash, int(totalInt)); err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "合并分片出错")
		return
	}

	// 保存到数据库
	err = dao.NewBackup().Add(config.Db, &model.Backup{
		BackupID:  common.RandomString(32, common.RandomLowercase|common.RandomMajuscule|common.RandomNumber),
		UserID:    userID,
		Name:      file.Filename,
		Game:      game,
		RealPath:  file7z,
		Size:      uint64(sizeInt),
		Hash:      hash,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Error(ctx, err.Error())
		common.ResponseErrorJSON(ctx, "保存数据错误")
		return
	}

	log.Info(ctx, "success")
	common.ResponseSuccess(ctx)
	return
}

// getPartPath 返回分片路径
func getPartPath(dir, hash string, total, index int) string {
	return fmt.Sprintf("%s/%s.%d.%d.part", dir, hash, total, index)
}
