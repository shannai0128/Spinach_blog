package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(val string) string {
	m :=md5.New() // 返回一个用于md5校验的接口
	m.Write([]byte(val)) // 写入要加密的数据
	return hex.EncodeToString(m.Sum(nil)) // 将加密后的byte转为string
}
