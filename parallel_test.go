package parallel

import (
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
