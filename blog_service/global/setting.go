package global

import (
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/logger"
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	JWTSetting      *setting.JWTSetting
	Logger          *logger.Logger
)
