package anime

import (
	"context"

	"github.com/vingarcia/ksql"
)

type Service struct {
	db ksql.Provider
}

func NewService(db ksql.Provider) Service {
	return Service{
		db: db,
	}
}

func (s Service) Create(ctx context.Context) error {
	return nil
}
