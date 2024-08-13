package list

import "sync"

func ConcurrentMap[T, U any](src []T, fn func(T) U) (dest []U) {
	size := len(src)
	dest = make([]U, size)
	dest = dest[:size] // this is safe

	var wg sync.WaitGroup
	for idx, item := range src {
		wg.Add(1)
		go func() {
			dest[idx] = fn(item)
			wg.Done()
		}()
	}

	wg.Wait()
	return
}

func FilterNil[T any](src []*T) []T {
	var dst []T

	for _, item := range src {
		if item != nil {
			dst = append(dst, *item)
		}
	}
	return dst
}
