package list_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ut-code/JourniCal/backend/pkg/list"
)

func TestConcurrentMap(t *testing.T) {
	assert := assert.New(t)

	src := []int{1, 2, 3, 4, 5}
	destch := WaitChan(func() []int {
		return list.ConcurrentMap(src, func(i int) int {
			r := rand.Intn(4) + 1
			// 1 <= sleep time < 5
			time.Sleep(time.Duration(r) * time.Millisecond)
			return 2 * i
		})
	})

	select {
	case <-time.After(5 * time.Millisecond):
		t.Error("timeout")
	case dest := <-destch:
		assert.Equal(dest, []int{2, 4, 6, 8, 10})
	}
}

func WaitChan[T any](fn func() T) <-chan T {
	var ch = make(chan T)
	go func() {
		ch <- fn()
	}()

	return ch
}
