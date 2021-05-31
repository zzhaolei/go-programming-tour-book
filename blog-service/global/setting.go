package global

import (
	"github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/logger"
	setting "github.com/zzhaolei/go-programming-tour-book/blog_service/pkg/settings"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	Logger *logger.Logger
)
