package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCar_Go_GoWithTwoCarsWithDifferentEngines_ShouldMakeDifferentSounds(t *testing.T) {
	// Arrange
	tEngine := NewToyotaEngine()
	hEngine := NewHondaEngine()

	toyota := NewCar(tEngine)
	honda := NewCar(hEngine)

	// Act
	sound1 := toyota.Go()
	sound2 := honda.Go()

	//Assert
	assert.Equal(t, "Toyota-a-a-a!", sound1)
	assert.Equal(t, "Honda-a-a-a!", sound2)
}
