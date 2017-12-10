package rand

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

func NextBytes(nByte int) []byte {
	r := make([]byte, nByte)
	_, err := rand.Read(r)
	if err != nil {
		panic(err)
	}

	return r
}

func NextHex(nByte int) string {
	return hex.Dump(NextBytes(nByte))
}

func NextBase64(nByte int) string {
	return base64.StdEncoding.EncodeToString(NextBytes(nByte))
}
