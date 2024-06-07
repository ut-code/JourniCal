package test

import "testing"

func TestTemplate(t *testing.T) {
	got := 1 + 1
	if got != 2 {
		t.Errorf("Test fail: Expected 1 + 1 to be equal to 2, got %d", got)
	}
}
