package mediator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcreteMediator_Notify(t *testing.T) {
	// Arrange
	farm := &Farm{}
	fabric := &Fabric{}
	shop := &Shop{}

	ConnectColleagues(farm, fabric, shop)

	// Act
	res1 := fabric.mediator.Notify("Farm")
	res2 := fabric.mediator.Notify("Shop")

	// Assert
	assert.Equal(t, "Tomatoes and potato!", res1)
	assert.Equal(t, "A lot of money!", res2)
}
