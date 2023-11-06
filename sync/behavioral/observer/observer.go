package observer

type Publisher interface {
	Attach(observer Observer)
	SetState(state string)
	Notify()
}

type Observer interface {
	Update(state string)
}

type ConcretePublisher struct {
	observers []Observer
	state     string
}

func NewPublisher() Publisher {
	return &ConcretePublisher{}
}

func (c *ConcretePublisher) Attach(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *ConcretePublisher) SetState(state string) {
	c.state = state
}

func (c *ConcretePublisher) Notify() {
	for _, e := range c.observers {
		e.Update(c.state)
	}
}

type ConcreteObserver struct {
	state string
}

func NewObserver() *ConcreteObserver {
	return &ConcreteObserver{}
}

func (c *ConcreteObserver) Update(state string) {
	c.state = state
}
