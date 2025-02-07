package internal

import (
	"Rope_Net/pkg/logger"
	"encoding/json"
	"os"
)

type DBConfig struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type Config struct {
	Database DBConfig `json:"database"`
}

func ReadDBConfig() (*DBConfig, error) {
	logger.Info("读取数据库配置文件")
	file, err := os.ReadFile("config/db_config.json")
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	var config Config
	logger.Info("反序列化JSON数据")
	err = json.Unmarshal(file, &config)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return &config.Database, nil
}
