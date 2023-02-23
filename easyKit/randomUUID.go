package easykit

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/google/uuid"
)

// 随机生成一个唯一的字符串uuid
func RandomUUID() string {
	return uuid.New().String()
}

// MD5加密32位
func Md5_32(data string) string {

	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5加密16位
func Md5_16(data string) string {
	return Md5_32(data)[8:24]
}
