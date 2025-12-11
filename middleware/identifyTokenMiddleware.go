package middleware

import (
	"Rope_Net/pkg/identify/token"
	"Rope_Net/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func IdentifyTokenMiddleware(c *gin.Context) {
	logger.Info("使用中间件获取Token")
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 10010,
			"info": "未提供授权信息",
		})
		c.Abort()
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
	// 调用验证函数
	user, isValid := token.IdentifyToken(tokenString)
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 10011,
			"info": "无效授权信息",
		})
		c.Abort()
	}
	c.Set("user", user)
	c.Next()
}
