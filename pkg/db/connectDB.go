package db

import (
	"Rope_Net/internal"
	"Rope_Net/pkg/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	logger.Info("连接数据库")
	dbConfig, err := internal.ReadDBConfig()
	if err != nil {
		logger.Error(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return db, nil
}
