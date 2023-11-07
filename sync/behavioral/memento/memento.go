package memento

type Originator struct {
	state string
}

func NewOriginator(state string) *Originator {
	return &Originator{state: state}
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) SetState(m *Memento) {
	o.state = m.GetState()
}

type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

type Caretaker struct {
	memento *Memento
}
