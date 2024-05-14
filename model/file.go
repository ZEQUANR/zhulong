package model

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var FileState = struct {
	ToBeReviewed    int // 待评审
	UnderReview     int // 评审中
	ReviewCompleted int // 评审完成
	Repulse         int // 打回
}{
	ToBeReviewed:    0,
	UnderReview:     1,
	ReviewCompleted: 2,
	Repulse:         3,
}

type File struct {
	Name string                // 文件名
	Path string                // 文件存储路径
	Blob *multipart.FileHeader // 文件本体
}

func (o *File) Save(c *gin.Context) error {
	if _, err := os.Stat(o.Path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("检查目录时出错: %w", err)
	}

	if err := os.MkdirAll(o.Path, 0777); err != nil {
		return fmt.Errorf("创建目录时出错: %w", err)
	}

	if err := c.SaveUploadedFile(o.Blob, path.Join(o.Path, o.Name)); err != nil {
		return fmt.Errorf("保存文件时出错: %w", err)
	}

	return nil
}

func (o *File) Delete() error {
	absPath, err := filepath.Abs(o.Path)
	if err != nil {
		return fmt.Errorf("无法获取绝对路径: %w", err)
	}

	info, err := os.Stat(absPath)
	if err != nil {
		return fmt.Errorf("无法获取路径信息: %w", err)
	}

	if !info.IsDir() {
		err = os.Remove(absPath)
		if err != nil {
			return fmt.Errorf("删除文件时出错: %w", err)
		}
	}

	return nil
}

func (o *File) GetPath() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to get current working directory: %v", err)
	}

	o.Path = path.Join(dir, "files")

	o.Path = path.Join(o.Path, time.Now().Format("2006-01-02"))

	o.Path = path.Join(o.Path, time.Now().Format("2006-01-02-15-04-05"))

	return nil
}

func (o *File) CheckPath() error {
	absPath, err := filepath.Abs(o.Path)
	if err != nil {
		return fmt.Errorf("无法获取绝对路径: %w", err)
	}

	if _, err := os.Stat(absPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("路径不存在: %w", err)
		}
		return fmt.Errorf("检查路径时出错: %w", err)
	}

	return nil
}
