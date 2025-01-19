package services

type IPingService interface {
	GetMessage() string
}

type PingService struct {
	Message string
}

func (ps *PingService) GetMessage() string {
	return "pong"
}
