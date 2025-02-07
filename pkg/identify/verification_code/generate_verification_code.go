package verification_code

import (
	"Rope_Net/pkg/logger"
	"math/rand"
	"time"
)

// GenerateVerificationCode 用于生成指定长度的数字验证码
func GenerateVerificationCode(length int) string {
	logger.Info("生成验证码")
	const charset = "0123456789"
	// 创建一个长度为 length 的字节切片，用于存储验证码
	code := make([]byte, length)
	// 使用当前时间的纳秒级 Unix 时间戳创建一个 rand.Source 对象
	source := rand.NewSource(time.Now().UnixNano())
	// 使用这个 rand.Source 对象创建一个新的随机数生成器
	r := rand.New(source)
	for i := range code {
		// 生成一个 0 到 len(charset) - 1 之间的随机整数
		index := r.Intn(len(charset))
		// 从字符集 charset 中选取对应索引的字符，填充到 code 切片中
		code[i] = charset[index]
	}
	// 将字节切片转换为字符串并返回
	return string(code)
}
