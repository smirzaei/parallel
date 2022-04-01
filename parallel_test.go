package parallel

import (
	"math"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestForEachToCallTheFunctionForEveryElement(t *testing.T) {
	// Arrange
	input := []int{1, 2, 3}
	expected := make([]int, 0, len(input))

	sampleFunc := func(x int) {
		expected = append(expected, x)
	}

	// Act
	ForEach(input, sampleFunc)

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

	sampleFunc1 := func(x int) {
		expected1 = append(expected1, x)
	}

	sampleFunc2 := func(x int) {
		expected2 = append(expected2, x)
	}

	sampleFunc3 := func(x int) {
		expected3 = append(expected3, x)
	}

	// Act
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		ForEach(input2, sampleFunc2)
		wg.Done()
	}()

	go func() {
		ForEach(input1, sampleFunc1)
		wg.Done()
	}()

	go func() {
		ForEach(input3, sampleFunc3)
		wg.Done()
	}()

	wg.Wait()

	// Assert
	assert.ElementsMatch(t, input1, expected1)
	assert.ElementsMatch(t, input2, expected2)
	assert.ElementsMatch(t, input3, expected3)
}

func TestForEachLimitToCallTheFunctionForEveryElementAndNotExceedTheLimit(t *testing.T) {
	// Arrange
	input := []int{1, 2, 3, 4, 5}
	expected := make([]int, 0, len(input))
	executionTime := make([]time.Time, 0, len(input))
	concurrencyLimit := 3

	sampleFunc := func(x int) {
		executionTime = append(executionTime, time.Now().UTC())
		expected = append(expected, x)
		time.Sleep(1 * time.Second)
	}

	// Act
	ForEachLimit(input, concurrencyLimit, sampleFunc)

	// Assert
	assert.ElementsMatch(t, input, expected)

	// Check each group has executing time close to one another
	checkTime := func(times []time.Time) {
		for i := 0; i < (len(times) - 1); i++ {
			t1 := times[i].Truncate(1 * time.Second)
			t2 := times[i+1].Truncate(1 * time.Second)

			assert.True(t, t1.Equal(t2))
		}
	}

	group1 := executionTime[0:3]
	checkTime(group1)

	group2 := executionTime[3:5]
	checkTime(group2)

	firstGroupExecutionTime := group1[0].Truncate(1 * time.Second)
	secondGroupExecutionTime := group2[0].Truncate(1 * time.Second)

	assert.True(t, firstGroupExecutionTime.Before(secondGroupExecutionTime))

	// for i := 0; i < len(input); i += concurrencyLimit {
	// 	from := i
	// 	to := int(math.Min(float64(len(executionTime)), float64(concurrencyLimit+i)))

	// 	group := executionTime[from:to]
	// 	for j := 0; j < (len(group) - 1); j++ {
	// 		t1 := executionTime[i+j].Truncate(1 * time.Second)
	// 		t2 := executionTime[i+j+1].Truncate(1 * time.Second)

	// 		assert.True(t, t1.Equal(t2))
	// 	}
	// }
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

func TestMapToWorkFineConcurrently(t *testing.T) {
	// Arrange
	input := []int{1, 2, 3, 4, 5}
	expected1 := []int{2, 4, 6, 8, 10}
	expected2 := []int{3, 6, 9, 12, 15}
	expected3 := []int{1, 4, 9, 16, 25}
	expected4 := []int{1, 8, 27, 64, 125}

	doubleFn := func(x int) int {
		return x * 2
	}

	tripleFn := func(x int) int {
		return x * 3
	}

	squareFn := func(x int) int {
		return int(math.Pow(float64(x), 2))
	}

	cubeFn := func(x int) int {
		return int(math.Pow(float64(x), 3))
	}

	// Act & Assert
	wg := &sync.WaitGroup{}
	wg.Add(4)

	go func() {
		result := Map(input, doubleFn)
		assert.Equal(t, expected1, result)
		wg.Done()
	}()

	go func() {
		result := Map(input, tripleFn)
		assert.Equal(t, expected2, result)
		wg.Done()
	}()

	go func() {
		result := Map(input, squareFn)
		assert.Equal(t, expected3, result)
		wg.Done()
	}()

	go func() {
		result := Map(input, cubeFn)
		assert.Equal(t, expected4, result)
		wg.Done()
	}()

	wg.Wait()
}
