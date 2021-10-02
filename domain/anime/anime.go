package anime

import "context"

type Service struct{}

func NewService() Service {
	return Service{}
}

func (s Service) Create(ctx context.Context) error {
	return nil
}
