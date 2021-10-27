package domain

import "context"

// AnimeService ...
type AnimeService interface {
	Create(ctx context.Context, anime Anime) error
	Update(ctx context.Context, anime Anime) error
}
