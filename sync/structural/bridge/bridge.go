package bridge

type Transport interface {
	Go() string
}

type Car struct {
	engine Engine
}

func NewCar(engine Engine) *Car {
	return &Car{engine: engine}
}

func (c *Car) Go() string {
	return c.engine.TurnOn()
}

type Engine interface {
	TurnOn() string
}

type HondaEngine struct{}

func NewHondaEngine() *HondaEngine {
	return &HondaEngine{}
}

func (h *HondaEngine) TurnOn() string {
	return "Honda-a-a-a!"
}

type ToyotaEngine struct{}

func NewToyotaEngine() *ToyotaEngine {
	return &ToyotaEngine{}
}

func (t *ToyotaEngine) TurnOn() string {
	return "Toyota-a-a-a!"
}
