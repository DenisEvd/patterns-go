package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContext_Sort(t *testing.T) {
	// Arrange
	arr1 := []int{5, 1, 4, 3, 2}
	arr2 := []int{5, 1, 4, 3, 2}

	sort1 := &OldSort{}
	sort2 := &InsertionSort{}

	ctx := Context{}

	// Act
	ctx.Algorithm(sort1)
	ctx.Sort(arr1)

	ctx.Algorithm(sort2)
	ctx.Sort(arr2)

	// Assert
	assert.Equal(t, []int{1, 2, 3, 4, 5}, arr1, arr2)
}
