package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic("无法生成随机字符串")
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
