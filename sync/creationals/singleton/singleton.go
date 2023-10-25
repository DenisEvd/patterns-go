package singleton

import "sync"

type Singleton struct {
	id int
}

var (
	instance *Singleton
	once     sync.Once
)

func NewInstance(id int) *Singleton {
	once.Do(func() {
		instance = &Singleton{id: id}
	})

	return instance
}
