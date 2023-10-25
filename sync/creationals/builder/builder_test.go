package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductBuilder_WithHeader_AddHeader_NoError(t *testing.T) {
	// Arrange
	prodBuilder := NewProductBuilder()

	// Act
	err := prodBuilder.WithHeader("head")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "head", prodBuilder.product.Header)
}

func TestProductBuilder_WithHeader_AddEmptyHeader_Error(t *testing.T) {
	// Arrange
	prodBuilder := NewProductBuilder()

	// Act
	err := prodBuilder.WithHeader("")

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrEmptyArgument)
	assert.Equal(t, "", prodBuilder.product.Header)
}

func TestProductBuilder_WithText_AddText_ShouldAddText(t *testing.T) {
	// Arrange
	prodBuilder := NewProductBuilder()

	// Act
	prodBuilder.WithText("text")

	// Assert
	assert.Equal(t, "text", prodBuilder.product.Text)
}

func TestProductBuilder_WithFooter_AddFooter_NoError(t *testing.T) {
	// Arrange
	prodBuilder := NewProductBuilder()

	// Act
	err := prodBuilder.WithFooter("foot")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "foot", prodBuilder.product.Footer)
}

func TestProductBuilder_WithFooter_AddEmptyFooter_Error(t *testing.T) {
	// Arrange
	prodBuilder := NewProductBuilder()

	// Act
	err := prodBuilder.WithFooter("")

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrEmptyArgument)
	assert.Equal(t, "", prodBuilder.product.Footer)
}
