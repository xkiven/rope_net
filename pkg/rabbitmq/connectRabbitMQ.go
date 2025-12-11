package rabbitmq

import (
	"Rope_Net/pkg/logger"
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	// RabbitMQ 连接地址
	url := "amqp://guest:guest@localhost:5672/"
	//连接到RabbitMQ服务器
	logger.Info("连接RabbitMQ服务器")
	conn, err := amqp.Dial(url)
	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}
	//创建一个通道
	logger.Info("创建一个通道")
	ch, err := conn.Channel()
	if err != nil {
		logger.Error(err)
		return nil, nil, err
	}
	return conn, ch, nil

}
