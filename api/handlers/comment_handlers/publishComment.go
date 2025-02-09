package comment_handlers

import (
	"Rope_Net/models"
	"Rope_Net/pkg/logger"
	"Rope_Net/pkg/rabbitmq"
	"encoding/json"
	"github.com/streadway/amqp"
)

func PublishComment(comment models.PostComment) error {

	//连接rabbitMQ
	conn, ch, err := rabbitmq.ConnectRabbitMQ()
	if err != nil {
		logger.Error("无法连接到rabbitmq", err)
		return err
	}
	defer conn.Close()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"PostComment_queue", //队列名称
		false,               //是否持久化
		false,               //是否自动删除
		false,               //是否排他
		false,               //是否等待
		nil,                 //额外参数
	)
	if err != nil {
		logger.Error("无法声明队列：", err)
	}

	commentBytes, err := json.Marshal(comment)
	if err != nil {
		logger.Error("无法序列化：", err)
		return err
	}
	err = ch.Publish(
		"",     //交换器
		q.Name, //路由键，使用队列名称
		false,  //强制
		false,  //立即
		amqp.Publishing{
			ContentType: "application/json",
			Body:        commentBytes,
		})
	if err != nil {
		logger.Error("无法发布评论", err)
		return err
	}
	return nil

}
