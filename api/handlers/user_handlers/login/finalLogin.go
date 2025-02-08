package login

import (
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func FinalLogin(c *gin.Context) {
	var input struct {
		VerificationCode string `json:"verificationCode"`
	}

	// 绑定请求参数
	logger.Info("绑定参数")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10001,
			"error":  "绑定参数错误",
		})
		logger.Error(err)
		return
	}

	// 从会话中获取用户名
	logger.Info("从会话中获取用户名")
	session := sessions.Default(c)
	username := session.Get("username")
	if username == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10008,
			"error":  "未找到用户会话信息",
		})
		return
	}

	// 从 go-cache 中获取验证码
	logger.Info("从 go-cache 中获取验证码")
	storedCode, found := verificationCodeCache.Get(username.(string))
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10009,
			"error":  "验证码已过期或未获取",
		})
		return
	}

	if storedCode.(string) != input.VerificationCode {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10011,
			"error":  "验证码错误",
		})
		return
	}

	// 验证码验证通过，从 go-cache 中删除已使用的验证码
	verificationCodeCache.Delete(username.(string))

	//生成token
	logger.Info("生成随机token")
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	tokenBytes := make([]byte, 16)
	for i := range tokenBytes {
		tokenBytes[i] = charset[r.Intn(len(charset))]
	}
	token := string(tokenBytes)

	//连接数据库
	db, err := db2.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10002,
			"error":  "连接数据库错误",
		})
		logger.Error(err)
		return
	}
	defer db2.CloseDB(db)
	//插入token
	logger.Info("更新用户token")
	result := db.Table("users").Where("username = ?", username.(string)).Update("token", token)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"info":   "插入数据失败",
		})
		logger.Error(result.Error)
		return
	}
	//将token传给客户端
	logger.Info("将token传给客户端")
	c.Writer.Header().Set("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "已成功登录",
	})

}
