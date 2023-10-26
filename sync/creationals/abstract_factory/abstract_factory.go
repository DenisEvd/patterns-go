package abstract_factory

type AbstractFactory interface {
	CreateArcher() Archer
	CreateWarrior() Warrior
}

type FireUnitFactory struct {
}

func NewFireUnitFactory() *FireUnitFactory {
	return &FireUnitFactory{}
}

func (f *FireUnitFactory) CreateArcher() Archer {
	return &FireArcher{}
}

func (f *FireUnitFactory) CreateWarrior() Warrior {
	return &FireWarrior{}
}

type IceUnitFactory struct {
}

func NewIceUnitFactory() *IceUnitFactory {
	return &IceUnitFactory{}
}

func (i *IceUnitFactory) CreateArcher() Archer {
	return &IceArcher{}
}

func (i *IceUnitFactory) CreateWarrior() Warrior {
	return &IceWarrior{}
}

type Archer interface {
	Shoot() int
}

type FireArcher struct {
}

func (f *FireArcher) Shoot() int {
	return 5
}

type IceArcher struct {
}

func (i *IceArcher) Shoot() int {
	return 10
}

type Warrior interface {
	Hit() int
}

type FireWarrior struct {
}

func (f *FireWarrior) Hit() int {
	return 5
}

type IceWarrior struct {
}

func (i *IceWarrior) Hit() int {
	return 10
}
