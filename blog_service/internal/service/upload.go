package service

import (
	"errors"
	"mime/multipart"
	"os"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/global"

	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/upload"
)

type FileInfo struct {
	Name      string
	AccessURL string
}

func UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	// 获取hash后的文件名称
	fileName := upload.GetFileName(fileHeader.Filename)
	uploadSavePath := upload.GetSavePath()
	dst := uploadSavePath + "/" + fileName
	// 检测文件类型是否支持
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}

	// 检测存储路径是否存在，不存在就创建
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}

	// 检测文件大小是否超出限制
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}

	// 检测上传目录是否具有权限
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permission")
	}
	// 存储文件
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessURL := global.AppSetting.UploadServerURL + "/" + fileName
	return &FileInfo{
		Name:      fileName,
		AccessURL: accessURL,
	}, nil
}
