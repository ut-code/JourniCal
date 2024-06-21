package hash_test

import (
	"github.com/ut-code/JourniCal/backend/pkg/hash"
	"github.com/ut-code/JourniCal/backend/pkg/test/assertion"
	"testing"
)

func TestSHA256(t *testing.T) {
	assert := assertion.New(t)
	helloHash, err := hash.SHA256("Hello, World!")
	assert.Nil(err)
	// this value is obtained via `echo -n '["Hello, World!"]' | sha256sum`
	assert.Eq(helloHash.Hex(), "485a704605793d42cd6d0289e0df983371fe04cf1d0be2948924308389029b04")
}
