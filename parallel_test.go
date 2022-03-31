package parallel

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEachToCallTheFunctionForEveryElement(t *testing.T) {
	// Arrange
	input := []int{1, 2, 3}
	expected := make([]int, 0, len(input))

	sampleFun := func(x int) {
		expected = append(expected, x)
	}

	// Act
	ForEach(input, sampleFun)

	// Assert
	assert.ElementsMatch(t, input, expected)
}

func TestForEachToWorkFineConcurrently(t *testing.T) {
	// Arrange
	input1 := []int{1, 2, 3}
	expected1 := make([]int, 0, len(input1))

	input2 := []int{4, 5, 6}
	expected2 := make([]int, 0, len(input2))

	input3 := []int{7, 8, 9}
	expected3 := make([]int, 0, len(input3))

	sampleFun1 := func(x int) {
		expected1 = append(expected1, x)
	}

	sampleFun2 := func(x int) {
		expected2 = append(expected2, x)
	}

	sampleFun3 := func(x int) {
		expected3 = append(expected3, x)
	}

	// Act
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		ForEach(input2, sampleFun2)
		wg.Done()
	}()

	go func() {
		ForEach(input1, sampleFun1)
		wg.Done()
	}()

	go func() {
		ForEach(input3, sampleFun3)
		wg.Done()
	}()

	wg.Wait()

	// Assert
	assert.ElementsMatch(t, input1, expected1)
	assert.ElementsMatch(t, input2, expected2)
	assert.ElementsMatch(t, input3, expected3)
}

func TestMapToCallTheFunctionForEveryElementAndReturnTheExpectedResult(t *testing.T) {
	// Arrange
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	doubleFn := func(x int) int {
		return x * 2
	}

	// Act
	result := Map(input, doubleFn)

	// Assert
	assert.Equal(t, expected, result)
}
