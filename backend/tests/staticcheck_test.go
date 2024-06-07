package test

import (
	"testing"
)

func TestStaticcheck(t *testing.T) {
	err := Error()
	UseLater(err) // remove this line and staticcheck will fail
	err = Error()
	UseLater(err)
}

func Error() error {
	return nil
}
func UseLater(_ ...any) {
}
