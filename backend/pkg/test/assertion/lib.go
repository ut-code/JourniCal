package assertion

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

type asserter struct {
	t *testing.T
}

func New(t *testing.T) asserter {
	return asserter{t: t}
}

func (a asserter) Error(m ...string) {
	a.t.Error("Assertion failed: ", strings.Join(m, " "))
}

func (a asserter) Be(b bool, m ...string) {
	if !b {
		a.Error(m...)
	}
}
func (a asserter) Eq(x, y any, m ...string) {
	if x != y {
		a.Error(strings.Join([]string{"Assert Eq failed: expected ", fmt.Sprint(y), ", Got ", fmt.Sprint(x), strings.Join(m, " ")}, ""))
	}
}
func (a asserter) NEq(x, y any, m ...string) {
	if x == y {
		a.Error(strings.Join([]string{"Assert NEq failed: right hand: ", fmt.Sprint(x), ", left hand: ", fmt.Sprint(y), strings.Join(m, " ")}, ""))
	}
}
func (a asserter) Not(b bool, m ...string) {
	if b {
		a.Error(m...)
	}
}
func (a asserter) Nil(e error, m ...string) {
	if e != nil {
		a.Error(strings.Join([]string{
			"Assert.Nil failed. Got: ",
			e.Error(),
			strings.Join(m, " "),
		}, ""))
	}
}

func (a asserter) NotNil(e error, m ...string) {
	if e == nil {
		a.Error(m...)
	}
}

func (a asserter) PanicOn(e error) {
	if e != nil {
		log.Fatalln("PanicOn called with non-nil error: ", e)
	}
}
