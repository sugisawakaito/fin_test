package service

import (
	"io"
	"crypto/sha1"
	"encoding/hex"
)

func GetEncriptedUserInfo(name, password string) string {
	info := name + password
	sha1 := sha1.New()
	io.WriteString(sha1, info)
	return hex.EncodeToString(sha1.Sum(nil))
}
