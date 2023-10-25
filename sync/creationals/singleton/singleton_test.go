package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleton_NewInstance_CreateInstanceTwice_ShouldCreateOnce(t *testing.T) {
	// Arrange
	ins1 := NewInstance(123)

	// Act
	ins2 := NewInstance(456)

	// Assert
	assert.Equal(t, 123, ins1.id, ins2.id)
}
