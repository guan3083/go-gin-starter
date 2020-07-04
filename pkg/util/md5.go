package util

import (
	"crypto/md5"
	"encoding/hex"
	"go-gin-starter/pkg/setting"
)

// EncodeMD5 md5 encryption
func EncodeMD5(value, salt string) string {
	m := md5.New()
	m.Write([]byte(value))
	m.Write([]byte(salt))
	return hex.EncodeToString(m.Sum(nil))
}

// 判断str的MD5加密是否为md5Str
func MD5Equals(str, salt, md5Str string) bool {
	return EncodeMD5(str, salt) == md5Str
}

func GetMd5String(str string) string {
	salt := setting.AppSetting.MD5Salt
	return EncodeMD5(str, salt)
}
