package db

import (
	"Rope_Net/pkg/logger"
	"gorm.io/gorm"
)

func CloseDB(db *gorm.DB) {
	logger.Info("确保在函数结束时关闭数据库连接")
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if err := sqlDB.Close(); err != nil {
		logger.Error(err.Error())
	}
	logger.Info("已断开数据库连接")

}
