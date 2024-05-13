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

type File struct {
	Name string
	Path string
	Blob *multipart.FileHeader
}

func (o *File) Save(c *gin.Context) error {
	if _, err := os.Stat(o.Path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("检查目录时出错: %w", err)
	}

	if err := os.MkdirAll(o.Path, 0777); err != nil {
		return fmt.Errorf("创建目录时出错: %w", err)
	}

	timeStr := time.Now().Format("2006年01月02日15时04分05秒-000000-")
	o.Path = path.Join(o.Path, timeStr+o.Name)

	if err := c.SaveUploadedFile(o.Blob, o.Path); err != nil {
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

	o.Path = dir

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
