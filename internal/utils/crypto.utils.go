package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func Hash(bv []byte) string {
	h := sha256.New()
	h.Write(bv)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
