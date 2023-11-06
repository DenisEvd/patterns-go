package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcretePublisher_Notify(t *testing.T) {
	// Arrange
	obs1 := NewObserver()
	obs2 := NewObserver()
	obs3 := NewObserver()

	publisher := NewPublisher()
	publisher.Attach(obs1)
	publisher.Attach(obs2)
	publisher.Attach(obs3)

	// Act
	publisher.SetState("Stable")
	publisher.Notify()

	// Assert
	assert.Equal(t, "Stable", obs1.state, obs2.state, obs3.state)
}
