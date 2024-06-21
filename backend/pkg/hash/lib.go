package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

type hash struct {
	b []byte
}

func SHA256(a ...any) (*hash, error) {
	j, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	b := sha256.Sum256(j)
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
