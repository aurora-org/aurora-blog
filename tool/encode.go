package tool

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 加密防止数据库存储明文密码
func Md5Encode(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
