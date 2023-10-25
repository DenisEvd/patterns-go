package prototype

type Prototype interface {
	Clone() Prototype
	GetName() string
}

func NewPrototype(name string) Prototype {
	return &Product{name: name}
}

type Product struct {
	name string
}

func (p *Product) Clone() Prototype {
	return &Product{name: p.name}
}

func (p *Product) GetName() string {
	return p.name
}
