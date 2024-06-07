package test

import "testing"

func TestTemplate(t *testing.T) {
	got := 1 + 1
	if got != 3 {
		t.Errorf("Test fail: Expected 1 + 1 to be equal to 3, got %d", got)
	}
}
