package main

import (
	"fmt"
	"time"

	"github.com/smirzaei/parallel"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7}

	parallel.ForEach(input, func(x int) {
		fmt.Printf("Processing %d\n", x)
	})

	fmt.Println("=============================")
	output := parallel.Map(input, func(x int) int {
		fmt.Printf("Processing %d\n", x)
		return x * 2
	})

	fmt.Printf("The final result is %v\n", output)

	fmt.Println("=============================")
	maxConcurrency := 2
	parallel.ForEachLimit(input, maxConcurrency, func(x int) {
		executionTime := time.Now().UTC().Format("15:04:05.999")
		fmt.Printf("%s - Processing %d\n", executionTime, x)

		time.Sleep(1 * time.Second)
	})

	fmt.Println("=============================")
	output2 := parallel.MapLimit(input, maxConcurrency, func(x int) int {
		executionTime := time.Now().UTC().Format("15:04:05.999")
		fmt.Printf("%s - Processing %d\n", executionTime, x)
		time.Sleep(1 * time.Second)

		return x * 2
	})

	fmt.Printf("The final result is %v\n", output2)
}
