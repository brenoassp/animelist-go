package domain

import "context"

// AnimeService ...
type AnimeService interface {
	Create(ctx context.Context) error
}
