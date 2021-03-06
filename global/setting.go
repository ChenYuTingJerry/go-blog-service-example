package global

import (
	"github/ChenYuTingJerry/blog-service/pkg/logger"
	"github/ChenYuTingJerry/blog-service/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger *logger.Logger
)
