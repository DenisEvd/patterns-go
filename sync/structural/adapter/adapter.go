package adapter

type Target interface {
	Request() string
}

type OldService struct{}

func (o *OldService) OldRequest() string {
	return "Request"
}

type ServiceAdapter struct {
	*OldService
}

func NewServiceAdapter(service *OldService) Target {
	return &ServiceAdapter{OldService: service}
}

func (s *ServiceAdapter) Request() string {
	return s.OldRequest()
}
