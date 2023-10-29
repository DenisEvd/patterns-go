package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComponent_Path_TryGetPath_ShouldReturnPath(t *testing.T) {
	// Arrange
	root := NewDirectory("C:")
	childOne := NewDirectory("OS")
	childTwo := NewFile("x64")

	err1 := root.Add(childOne)
	err2 := childOne.Add(childTwo)

	// Act
	path := root.Path("x64")

	// Assert
	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.Equal(t, "C:/OS/x64", path)
}

func TestComponent_Add_AddComponentToLeave_ShouldReturnError(t *testing.T) {
	// Arrange
	root := NewDirectory("C:")
	childOne := NewDirectory("OS")
	childTwo := NewFile("x64")
	childThree := NewFile("x86")

	_ = root.Add(childOne)
	_ = childOne.Add(childTwo)

	// Act
	err := childTwo.Add(childThree)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, ErrCanNotAddComponentToLeave, err)
}
