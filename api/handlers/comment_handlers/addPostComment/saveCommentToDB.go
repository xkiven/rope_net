package addPostComment

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
)

func SaveCommentToDB(comment models.PostComment) error {
	//连接数据库
	db, err := db2.ConnectDB()
	if err != nil {

		logger.Error(err)
		return err
	}
	defer db2.CloseDB(db)
	//查找此post
	logger.Info("检查此post是否存在")
	var post models.Post
	result := db.Where("id = ?", comment.PostID).First(&post)
	if result.Error != nil {

		logger.Error(result.Error)
		return err
	}
	//插入数据
	logger.Info("创建评论并存入数据库")
	result = db.Create(&comment)
	if result.Error != nil {

		logger.Error(result.Error)
		return err
	}
	return nil
	////连接rabbitMQ
	//conn, ch, err := rabbitmq.ConnectRabbitMQ()
	//if err != nil {
	//	logger.Error("无法连接到rabbitmq", err)
	//	return err
	//}
	//defer conn.Close()
	//defer ch.Close()
	//
	//q, err := ch.QueueDeclare(
	//	"PostComment_queue", //队列名称
	//	false,               //是否持久化
	//	false,               //是否自动删除
	//	false,               //是否排他
	//	false,               //是否等待
	//	nil,                 //额外参数
	//)
	//if err != nil {
	//	logger.Error("无法声明队列：", err)
	//	return err
	//}
	//
	//commentBytes, err := json.Marshal(comment)
	//if err != nil {
	//	logger.Error("无法序列化：", err)
	//	return err
	//}
	//err = ch.Publish(
	//	"",     //交换器
	//	q.Name, //路由键，使用队列名称
	//	false,  //强制
	//	false,  //立即
	//	amqp.Publishing{
	//		ContentType: "application/json",
	//		Body:        commentBytes,
	//	})
	//if err != nil {
	//	logger.Error("无法发布评论", err)
	//	return err
	//}
	//return nil
}
