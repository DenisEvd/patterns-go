package facade

import "strings"

type Man struct {
	house *House
	tree  *Tree
	child *Child
}

func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

func (m *Man) Life() string {
	strBuilder := strings.Builder{}

	strBuilder.WriteString(m.house.Do() + "\n")
	strBuilder.WriteString(m.tree.Do() + "\n")
	strBuilder.WriteString(m.child.Do() + "\n")

	return strBuilder.String()
}

type House struct{}

func (h *House) Do() string {
	return "Build house!"
}

type Tree struct{}

func (t *Tree) Do() string {
	return "Grow tree!"
}

type Child struct{}

func (c *Child) Do() string {
	return "Child born!"
}
