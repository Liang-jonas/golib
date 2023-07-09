package cipher

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5WithoutSalt(text string) string {
	return Md5WithSalt(text, "")
}

func Md5WithSalt(text, salt string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
