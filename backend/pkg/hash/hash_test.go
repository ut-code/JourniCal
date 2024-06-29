package hash_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ut-code/JourniCal/backend/pkg/hash"
	"testing"
)

func TestSHA256(t *testing.T) {
	assert := assert.New(t)
	// this hard-coded hash is obtained via `echo -n '"Hello, ""World!"' | sha256sum`
	assert.Equal(hash.SHA256("Hello, ", "World!").Hex(), "8e2c61419c2e3ad5e525f2c7045232ed054dc522ccf68ea00c7904d64d2753d2")
}
