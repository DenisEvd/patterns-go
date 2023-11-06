package visitor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCity_Accept_ShouldReturnListOfActions(t *testing.T) {
	// Arrange
	tourist := People{}

	city := City{}
	city.Add(&Zoo{})
	city.Add(&Market{})
	city.Add(&Park{})

	// Act
	res := city.Accept(&tourist)

	// Assert
	assert.Equal(t, "Feed lion\nBuy fruits\nMake photo in the park\n", res)
}
