package login

import (
	"Rope_Net/models"
	db2 "Rope_Net/pkg/db"
	"Rope_Net/pkg/identify/verification_code"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"net/http"
	"time"
)

// 初始化一个默认过期时间为 5 分钟，清理间隔为 10 分钟的缓存
var verificationCodeCache = cache.New(5*time.Minute, 10*time.Minute)

func PreLogin(c *gin.Context) {
	var user models.User
	logger.Info("绑定参数")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 10001,
			"error":  "绑定参数错误",
		})
		logger.Error(err)
		return
	}
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
	var existingUser models.User
	logger.Info("通过用户名查找")
	result := db.Where("username = ?", user.Username).First(&existingUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10003,
			"error":  "用户名错误",
		})
		logger.Error(err)
		return
	}

	logger.Info("比对密码")
	if existingUser.Password != user.Password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10004,
			"error":  "密码错误",
		})
		return
	}
	//生成验证码
	verificationCode := verification_code.GenerateVerificationCode(4)
	logger.Info("储存验证码到缓存")

	// 存储验证码到缓存
	verificationCodeCache.Set(user.Username, verificationCode, cache.DefaultExpiration)

	//使用gin会话管理储存用户名
	logger.Info("创建会话")
	session := sessions.Default(c)
	session.Set("username", user.Username)
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10005,
			"info":   "创建会话失败",
		})
		logger.Error(err)
		return
	}
	//发送验证码
	err = verification_code.SendVerificationCode(user.Email, verificationCode)
	if err != nil {
		logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 10006,
			"info":   "发送验证码失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 10000,
		"info":   "验证码已发送",
	})

}
