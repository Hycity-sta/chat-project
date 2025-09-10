package utils

import (
	"strings"

	"crypto/md5"
	"encoding/hex"
)

// 小写格式的md5
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写格式的md5
func MD5Encode_(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密：将明文密码 plainpwd和盐值salt拼接后计算 MD5 哈希
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// 解密
func ValidPassword(plainpwd, salt string, password string) bool {
	md := Md5Encode(plainpwd + salt)
	return md == password
}
