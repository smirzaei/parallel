package parallel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForEachToCallTheFunctionForEveryElement(t *testing.T) {
	// Arrange
	input := []int{1, 2, 3}
	output := make([]int, 0, len(input))

	sampleFun := func(x int) {
		output = append(output, x)
	}

	// Act
	ForEach(input, sampleFun)

	// Assert
	assert.ElementsMatch(t, input, output)
}
