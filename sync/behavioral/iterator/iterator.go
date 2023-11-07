package iterator

type Iterator interface {
	Value() interface{}
	Next()
	Prev()
	Reset()
	End()
}

type Aggregate interface {
	Iterator() Iterator
}

type BookShelf struct {
	books []Book
}

type BookShelfIterator struct {
	shelf *BookShelf
	idx   int
}

func (b *BookShelfIterator) Value() interface{} {
	return b.shelf.books[b.idx]
}

func (b *BookShelfIterator) Next() {
	b.idx++

	if b.idx == len(b.shelf.books) {
		b.idx = 0
	}
}

func (b *BookShelfIterator) Prev() {
	b.idx--

	if b.idx == -1 {
		b.idx = len(b.shelf.books) - 1
	}
}

func (b *BookShelfIterator) Reset() {
	b.idx = 0
}

func (b *BookShelfIterator) End() {
	b.idx = len(b.shelf.books) - 1
}

func (b *BookShelf) Iterator() Iterator {
	return &BookShelfIterator{
		shelf: b,
		idx:   0,
	}
}

func (b *BookShelf) Add(name string) {
	b.books = append(b.books, Book{name: name})
}

type Book struct {
	name string
}
