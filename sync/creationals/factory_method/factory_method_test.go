package factory_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductFactory_Create_TakeA_ShouldCreateProductA(t *testing.T) {
	// Arrange
	fact := NewProductFactory()

	// Act
	prod := fact.Create("A")

	// Assert

	assert.Equal(t, &ProductA{}, prod)
}

func TestProductFactory_Create_TakeB_ShouldCreateProductB(t *testing.T) {
	// Arrange
	fact := NewProductFactory()

	// Act
	prod := fact.Create("B")

	// Assert

	assert.Equal(t, &ProductB{}, prod)
}
