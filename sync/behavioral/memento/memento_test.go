package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemento_GetState(t *testing.T) {
	// Arrange
	orig := NewOriginator("Stable")
	care := Caretaker{memento: orig.CreateMemento()}

	// Act
	orig.state = "asd13"
	orig.SetState(care.memento)

	// Assert
	assert.Equal(t, "Stable", orig.state)
}
