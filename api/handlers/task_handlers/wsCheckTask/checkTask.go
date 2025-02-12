package wsCheckTask

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"fmt"
	"time"
)

func CheckTask() {
	now := time.Now()
	fmt.Printf("当前时间: %s，时区: %s\n", now.Format(time.RFC3339), now.Location())
	//连接数据库
	db, err := db2.ConnectDB()
	if err != nil {
		logger.Error(err)
		return
	}
	defer db2.CloseDB(db)
	//查找为未完成的任务
	logger.Info("每分钟检查一次")
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		var tasks []models.Task
		db = db.Debug()
		result := db.Where("completed = ? AND deadline < ?", false, time.Now().Format(time.RFC3339)).Find(&tasks)
		if result.Error != nil {
			logger.Error(result.Error)
			continue
		}

		for _, task := range tasks {
			reminder := fmt.Sprintf("任务提醒: 任务 '%s' (ID: %d) 已到截止时间，但尚未完成。", task.Name, task.ID)
			logger.Info(reminder)
			// 发送提醒消息给对应用户的客户端
			SendRemind(task.UserID, reminder)

		}

	}

}
