package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProduct_Clone(t *testing.T) {
	// Arrange
	product := NewPrototype("Car")

	// Act
	productClone := product.Clone()

	// Assert
	assert.Equal(t, "Car", productClone.GetName())
}
