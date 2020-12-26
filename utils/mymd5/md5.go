package mymd5

import (
	"crypto/md5"
	"encoding/hex"
)

func Encryption(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
