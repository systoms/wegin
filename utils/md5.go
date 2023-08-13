package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func CalculateMD5(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
