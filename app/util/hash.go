package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(s string) (hash string) {
	m := md5.New()
	s += AppSetting.Md5String
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
