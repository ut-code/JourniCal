package hash_test

import (
	"github.com/ut-code/JourniCal/backend/pkg/hash"
	"github.com/ut-code/JourniCal/backend/pkg/test/assertion"
	"testing"
)

func TestSHA256(t *testing.T) {
	assert := assertion.New(t)
	// this hard-coded hash is obtained via `echo -n '"Hello, ""World!"' | sha256sum`
	assert.Eq(hash.SHA256("Hello, ", "World!").Hex(), "8e2c61419c2e3ad5e525f2c7045232ed054dc522ccf68ea00c7904d64d2753d2")
}
