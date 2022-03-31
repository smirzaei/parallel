package parallel

import (
	"sync"
)

func ForEach[T any](arr []T, fn func(T)) {
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))

	for _, item := range arr {
		go func(x T) {
			fn(x)
			wg.Done()
		}(item)
	}

	wg.Wait()
}

func Map[T1 any, T2 any](arr []T1, fn func(T1) T2) []T2 {
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))

	output := make([]T2, len(arr), len(arr))

	for i := range arr {
		go func(index int, x T1) {
			result := fn(x)
			output[index] = result

			wg.Done()
		}(i, arr[i])
	}

	wg.Wait()
	return output
}
