package factory_method

type Factory interface {
	Create(name string) Product
}

type ProductFactory struct {
}

func NewProductFactory() *ProductFactory {
	return &ProductFactory{}
}

func (p *ProductFactory) Create(name string) Product {
	switch name {
	case "A":
		return &ProductA{}
	default:
		return &ProductB{}
	}
}

type Product interface {
	GetName() string
}

type ProductA struct {
	name string
}

func (p *ProductA) GetName() string {
	return p.name
}

type ProductB struct {
	name string
}

func (p *ProductB) GetName() string {
	return p.name
}
