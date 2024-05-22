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
	ToBeUploaded    int // 待上传
	ToBeReviewed    int // 待评审
	UnderReview     int // 评审中
	ReviewCompleted int // 评审完成
	Repulse         int // 打回
	Deprecation     int // 弃用
}{
	ToBeUploaded:    1,
	ToBeReviewed:    2,
	UnderReview:     3,
	ReviewCompleted: 4,
	Repulse:         5,
	Deprecation:     6,
}

type File struct {
	Name string                // 文件名
	Path string                // 文件存储路径
	Blob *multipart.FileHeader // 文件本体
}

func (o *File) Save(c *gin.Context) error {
	if _, err := os.Stat(o.Path); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("struct: File | func: Save | index: 0 | err: Error checking directory: %w", err)
	}

	if err := os.MkdirAll(o.Path, 0777); err != nil {
		return fmt.Errorf("struct: File | func: Save | index: 1 | err: Error creating directory: %w", err)
	}

	if err := c.SaveUploadedFile(o.Blob, path.Join(o.Path, o.Name)); err != nil {
		return fmt.Errorf("struct: File | func: Save | index: 2 | err: Error saving file: %w", err)
	}

	return nil
}

func (o *File) Delete() error {
	absPath, err := filepath.Abs(o.Path)
	if err != nil {
		return fmt.Errorf("struct: File | func: Delete | index: 0 | err: Cannot get absolute path: %w", err)
	}

	info, err := os.Stat(absPath)
	if err != nil {
		return fmt.Errorf("struct: File | func: Delete | index: 1 | err: Failed to obtain path information: %w", err)
	}

	if !info.IsDir() {
		err = os.Remove(absPath)
		if err != nil {
			return fmt.Errorf("struct: File | func: Delete | index: 2 | err: Error deleting file: %w", err)
		}
	}

	return nil
}

func (o *File) GetPath() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("struct: File | func: GetPath | index: 0 | err: Unable to get current working directory: %w", err)
	}

	o.Path = path.Join(dir, "files")
	o.Path = path.Join(o.Path, time.Now().Format("2006-01-02"))
	o.Path = path.Join(o.Path, time.Now().Format("2006-01-02-15-04-05"))

	return nil
}

func (o *File) CheckPath() error {
	absPath, err := filepath.Abs(o.Path)
	if err != nil {
		return fmt.Errorf("struct: File | func: CheckPath | index: 0 | err: Cannot get absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("struct: File | func: CheckPath | index: 1 | err: Path does not exist: %w", err)
		}
		return fmt.Errorf("struct: File | func: CheckPath | index: 2 | err: Error checking path: %w", err)
	}

	return nil
}
