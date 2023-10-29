package proxy

import "log"

type Service interface {
	Handle(x int) int
}

type ServiceProxy struct {
	service Service
}

func NewServiceProxy(service Service) *ServiceProxy {
	return &ServiceProxy{service: service}
}

func (s *ServiceProxy) Handle(x int) int {
	log.Printf("Service got: %d", x)

	return s.service.Handle(x)
}

type ConcreteService struct{}

func NewConcreteService() *ConcreteService {
	return &ConcreteService{}
}

func (c *ConcreteService) Handle(x int) int {
	return x * x
}
