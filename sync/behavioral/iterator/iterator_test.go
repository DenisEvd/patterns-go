package iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookShelf_Iterator(t *testing.T) {
	// Arrange
	shelf := &BookShelf{}
	shelf.Add("1")
	shelf.Add("2")
	shelf.Add("10")

	it := shelf.Iterator()

	// Act
	res1 := it.Value().(Book).name
	it.Next()
	res2 := it.Value().(Book).name
	it.Next()
	res3 := it.Value().(Book).name

	// Assert
	assert.Equal(t, "1", res1)
	assert.Equal(t, "2", res2)
	assert.Equal(t, "10", res3)
}
