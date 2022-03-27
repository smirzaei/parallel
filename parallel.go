package parallel

import "sync"

func ForEach[T any](arr []T, fn func(T)) {
	wg := new(sync.WaitGroup)
	wg.Add(len(arr))

	for _, item := range arr {
		go func(x T) {
			fn(x)
			wg.Done()
		}(item)
	}

	wg.Wait()
}
