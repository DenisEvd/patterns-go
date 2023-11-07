package mediator

type Mediator interface {
	Notify(msg string) string
}

type ConcreteMediator struct {
	farm   *Farm
	fabric *Fabric
	shop   *Shop
}

func (c *ConcreteMediator) Notify(msg string) string {
	switch msg {
	case "Farm":
		return c.farm.GetVegetables()
	case "Fabric":
		return c.fabric.Fetch()
	case "Shop":
		return c.shop.Sell()
	default:
		return ""
	}
}

func ConnectColleagues(farm *Farm, fabric *Fabric, shop *Shop) {
	mediator := &ConcreteMediator{
		farm:   farm,
		fabric: fabric,
		shop:   shop,
	}

	farm.mediator = mediator
	fabric.mediator = mediator
	shop.mediator = mediator
}

type Farm struct {
	mediator Mediator
}

func (f *Farm) GetVegetables() string {
	return "Tomatoes and potato!"
}

type Fabric struct {
	mediator Mediator
}

func (f *Fabric) Fetch() string {
	return "Chips and juice!"
}

type Shop struct {
	mediator Mediator
}

func (s *Shop) Sell() string {
	return "A lot of money!"
}
