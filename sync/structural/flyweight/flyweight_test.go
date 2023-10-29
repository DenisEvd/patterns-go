package flyweight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcreteImage_Draw(t *testing.T) {
	// Arrange
	factory := NewImageFlyweightFactory()
	image := factory.GetImage("kitty")

	// Act
	res := image.Draw(100, 200)

	// Assert
	assert.Equal(t, "Draw image: kitty\nWidth: 100 Height: 200", res)
}
