package internal

import (
	"Rope_Net/pkg/logger"
	"encoding/json"
	"os"
)

type QQEmailConfig struct {
	QQEmail         string `json:"qq_email"`
	QQEmailAuthCode string `json:"qq_email_auth_code"`
	SMTPServer      string `json:"smtp_server"`
	SMTPPort        string `json:"smtp_port"`
}

func ReadQQEmailConfig() (*QQEmailConfig, error) {
	logger.Info("读取qq邮箱配置文件")
	file, err := os.ReadFile("config/qq_email_config.json")
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var qqEmailConfig QQEmailConfig
	err = json.Unmarshal(file, &qqEmailConfig)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &qqEmailConfig, nil
}
