package rand

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strconv"
)

func NextBytes(nBytes int) []byte {
	r := make([]byte, nBytes)
	_, err := rand.Read(r)
	if err != nil {
		panic(err)
	}

	return r
}

func NextHex(nBytes int) string {
	return hex.EncodeToString(NextBytes(nBytes))
}

func NextNumberFixedLength(nLength int) string {
	b := NextBytes(nLength)
	buf := bytes.NewBufferString("")
	for _, v := range b {
		buf.WriteString(strconv.Itoa(int(v) % 10))
	}
	return buf.String()
}

func NextNumber(nBytes int) string {
	b := NextBytes(nBytes)
	return (&big.Int{}).SetBytes(b).Text(10)
}
