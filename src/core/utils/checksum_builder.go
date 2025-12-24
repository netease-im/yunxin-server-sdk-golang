package utils

import (
	"crypto/sha1"
	"fmt"
)

// CheckSumBuilder 校验和构建器
type CheckSumBuilder struct{}

// HEX_DIGITS 十六进制数字
var HEX_DIGITS = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

// GetCheckSum 获取校验和
func GetCheckSum(appSecret, nonce, curTime string) string {
	return encode("sha1", appSecret+nonce+curTime)
}

// encode 编码
func encode(algorithm, value string) string {
	if value == "" {
		return ""
	}

	switch algorithm {
	case "sha1":
		hash := sha1.Sum([]byte(value))
		return getFormattedText(hash[:])
	default:
		panic(fmt.Sprintf("Unsupported algorithm: %s", algorithm))
	}
}

// getFormattedText 格式化文本
func getFormattedText(bytes []byte) string {
	length := len(bytes)
	buf := make([]byte, length*2)

	for i, b := range bytes {
		buf[i*2] = HEX_DIGITS[(b>>4)&0x0f]
		buf[i*2+1] = HEX_DIGITS[b&0x0f]
	}

	return string(buf)
}
