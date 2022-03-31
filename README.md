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


License
===

[MIT](https://github.com/smirzaei/parallel/blob/master/LICENSE)
