package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractFactory_CreateArchers_ShouldHaveDifferentDamage(t *testing.T) {
	// Arrange
	var factory1 AbstractFactory = NewIceUnitFactory()
	var factory2 AbstractFactory = NewFireUnitFactory()

	// Act
	arch1, arch2 := factory1.CreateArcher(), factory2.CreateArcher()

	// Assert
	assert.Equal(t, 10, arch1.Shoot())
	assert.Equal(t, 5, arch2.Shoot())
}
