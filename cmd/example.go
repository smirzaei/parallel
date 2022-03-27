package main

import (
	"fmt"

	"github.com/smirzaei/parallel"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6}

	parallel.ForEach(input, func(x int) {
		fmt.Printf("Processing %d\n", x)
	})

	fmt.Println("=============================")
	output := parallel.Map(input, func(x int) int {
		fmt.Printf("Processing %d\n", x)
		return x * 2
	})

	fmt.Printf("The final result is %v\n", output)
}
