package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
