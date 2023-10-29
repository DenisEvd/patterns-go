package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceDecorator_SomeMethod_ShouldReturnNewBehaviour(t *testing.T) {
	// Arrange
	service := NewService()
	serviceDecorator := NewServiceDecorator(service)

	// Act
	res := serviceDecorator.SomeMethod()

	// Assert
	assert.Equal(t, "New action", res)
}
