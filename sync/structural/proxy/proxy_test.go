package proxy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceProxy_Handle(t *testing.T) {
	// Arrange
	service := NewConcreteService()
	proxy := NewServiceProxy(service)

	// Act
	x := proxy.Handle(5)

	//Assert
	assert.Equal(t, 25, x)
}
