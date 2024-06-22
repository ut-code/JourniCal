package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type hash struct {
	b []byte
}

func SHA256(a ...any) (*hash, error) {
	v := []byte{}
	for _, item := range a {
		j, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		v = append(v, j...)
	}
	fmt.Println(string(v))
	b := sha256.Sum256(v)
	return &hash{
		b: b[:],
	}, nil
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
