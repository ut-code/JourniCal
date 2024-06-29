package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"
)

type hash struct {
	b []byte
}

func SHA256(a ...any) hash {
	v := []byte{}
	for _, item := range a {
		j, err := json.Marshal(item)
		if err != nil {
			log.Fatalln("json.Marshal failed to be called on: ", item)
		}
		v = append(v, j...)
	}
	b := sha256.Sum256(v)
	return hash{
		b: b[:],
	}
}

func (h hash) Bytes() []byte {
	return h.b
}

func (h hash) Hex() string {
	return hex.EncodeToString(h.b)
}

func (h hash) Base64() string {
	return base64.StdEncoding.EncodeToString(h.b)
}
