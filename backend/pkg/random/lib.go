package random

import (
	"crypto/rand"
	"math/big"

	"github.com/ut-code/JourniCal/backend/pkg/helper"
)

var pool = []byte("abcdefghijklmnopqrstuwvxyzABCDEFGHIJKLMNOPQRSTUWVXYZ_")
var poolSize = int64(len(pool))

func String(len int) string {
	ret := []byte{}
	for range len {
		n, err := rand.Int(rand.Reader, big.NewInt(poolSize))
		helper.ErrorLog(err)
		ret = append(ret, pool[n.Int64()])
	}
	return string(ret)
}

func Int(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	helper.ErrorLog(err)
	return int(n.Int64())
}
