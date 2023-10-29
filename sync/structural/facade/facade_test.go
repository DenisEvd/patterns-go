package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMan_Life_ShouldReturnLifePlan(t *testing.T) {
	// Arrange
	man := NewMan()

	// Act
	plan := man.Life()

	// Assert
	assert.Equal(t, "Build house!\nGrow tree!\nChild born!\n", plan)
}
