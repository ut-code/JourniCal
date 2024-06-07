package test

import (
	"testing"
)

func TestStaticcheck(t *testing.T) {
	err := Error()
	UseLater(err)
	err = Error()
	UseLater(err)
}

func Error() error {
	return nil
}
func UseLater(_ ...any) {
}
