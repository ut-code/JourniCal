package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/ut-code/JourniCal/backend/app/env/options"
	_ "github.com/ut-code/JourniCal/backend/pkg/tests/run-test-at-root"
)

func TestMain(t *testing.T) {
	expect := "this env variable is for testing env package"
	got := os.Getenv("ENV_TEST")
	assert.Equal(t, expect, got)

	expect = "test with double quotes"
	got = os.Getenv("ENV_TEST2")
	assert.Equal(t, expect, got)
}
