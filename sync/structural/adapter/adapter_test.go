package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceAdapter_Request(t *testing.T) {
	// Arrange
	oldService := &OldService{}
	var service = NewServiceAdapter(oldService)

	// Act
	res := service.Request()

	// Assert
	assert.Equal(t, "Request", res)
}
