package token

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"time"
)

func IdentifyToken(token string) bool {
	logger.Info("验证token")
	//连接数据库
	db, err := db2.ConnectDB()
	if err != nil {
		logger.Error(err)
		return false
	}
	defer db2.CloseDB(db)
	//验证
	var existingUser models.User
	result := db.Where("token = ? AND token_expires_at >?", token, time.Now()).First(&existingUser)
	if result.Error != nil || result.RowsAffected == 0 {
		return false
	}

	return true
}
