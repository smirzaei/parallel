Go Parallel
===

![CI Status](https://github.com/smirzaei/parallel/actions/workflows/test.yml/badge.svg)

Run go loops in parallel.


Installation
---

Run

```BASH
go get -u github.com/smirzaei/parallel
```

Examples
---

### ForEach

Call the given function once for every element. *The execution order is random*

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

Call the given function once for ever element and return a new slice with its results. *The slice order is preserved but the execution order is random*

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
