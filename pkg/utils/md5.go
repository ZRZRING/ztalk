package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(context string) string {
	h := md5.New()
	h.Write([]byte(context))
	return hex.EncodeToString(h.Sum([]byte(context)))
}
