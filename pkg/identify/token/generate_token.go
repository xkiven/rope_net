package token

import (
	"Rope_Net/pkg/logger"
	"math/rand"
	"time"
)

func GenerateToken() string {
	//生成token
	logger.Info("生成随机token")
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	tokenBytes := make([]byte, 16)
	for i := range tokenBytes {
		tokenBytes[i] = charset[r.Intn(len(charset))]
	}
	return string(tokenBytes)
}
