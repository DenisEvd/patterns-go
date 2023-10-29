package decorator

import "strings"

type Component interface {
	SomeMethod() string
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SomeMethod() string {
	return "Action"
}

type ServiceDecorator struct {
	*Service
}

func NewServiceDecorator(service *Service) *ServiceDecorator {
	return &ServiceDecorator{Service: service}
}

func (s *ServiceDecorator) SomeMethod() string {
	return "New " + strings.ToLower(s.Service.SomeMethod())
}
