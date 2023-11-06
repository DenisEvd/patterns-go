package template_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuotes_Quotes_ShouldReturnWithBothQuotes(t *testing.T) {
	// Arrange
	french := &FrenchQuotes{}
	german := &GermanQuotes{}

	quotesF := NewQuotes(french)
	quotesG := NewQuotes(german)

	// Act
	resF := quotesF.Quotes("Hello")
	resG := quotesG.Quotes("Hello")

	// Assert
	assert.Equal(t, "«Hello»", resF)
	assert.Equal(t, "„Hello“", resG)
}
