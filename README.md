Go Parallel
===

![CI Status](https://github.com/smirzaei/parallel/actions/workflows/test.yml/badge.svg)

Run Go loops in parallel.


Installation
---

Run

```BASH
go get -u github.com/smirzaei/parallel
```

Examples
---

### ForEach

Call the given function once for every element. Please note that the execution order is random.

```GO
input := []int{1, 2, 3, 4, 5, 6}

parallel.ForEach(input, func(x int) {
  fmt.Printf("Processing %d\n", x)
})

// Output:

// Processing 6
// Processing 3
// Processing 4
// Processing 5
// Processing 1
// Processing 2
```

### Map

Call the given function once for every element and return a new slice with its results. Please note that the order of the elements of the output slice is the same as the input slice and it is preserved but the execution order is random.

```GO
input := []int{1, 2, 3, 4, 5, 6}

result := parallel.Map(input, func(x int) int {
  fmt.Printf("Processing %d\n", x)
  return x * 2
})

fmt.Printf("The final result is %v\n", result)

// Output:

// Processing 6
// Processing 1
// Processing 2
// Processing 3
// Processing 4
// Processing 5
// The final result is [2 4 6 8 10 12]
```

## Limiting the concurrency level

If you need to limit the number parallel executions, you can use the following functions:

### ForEachLimit

```GO
input := []int{1, 2, 3, 4, 5, 6, 7}

maxConcurrency := 2
parallel.ForEachLimit(input, maxConcurrency, func(x int) {
  executionTime := time.Now().UTC().Format("15:04:05.999")
  fmt.Printf("%s - Processing %d\n", executionTime, x)

  time.Sleep(1 * time.Second)
})

// Output:

// 09:47:54.071 - Processing 2
// 09:47:54.071 - Processing 1
// 09:47:55.071 - Processing 3
// 09:47:55.071 - Processing 4
// 09:47:56.071 - Processing 5
// 09:47:56.071 - Processing 6
// 09:47:57.071 - Processing 7
```

### MapLimit

```GO
input := []int{1, 2, 3, 4, 5, 6, 7}

maxConcurrency := 2
output := parallel.MapLimit(input, maxConcurrency, func(x int) int {
  executionTime := time.Now().UTC().Format("15:04:05.999")
  fmt.Printf("%s - Processing %d\n", executionTime, x)
  time.Sleep(1 * time.Second)

  return x * 2
})

fmt.Printf("The final result is %v\n", output)

// Output:

// 09:53:58.998 - Processing 2
// 09:53:58.998 - Processing 1
// 09:53:59.998 - Processing 3
// 09:53:59.998 - Processing 4
// 09:54:00.999 - Processing 5
// 09:54:00.999 - Processing 6
// 09:54:01.999 - Processing 7
// The final result is [2 4 6 8 10 12 14]
```

License
===

[MIT](https://github.com/smirzaei/parallel/blob/master/LICENSE)
