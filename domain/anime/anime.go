package anime

import (
	"context"

	"github.com/brenoassp/animelist-go/domain"
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

func (s Service) Create(ctx context.Context, anime domain.Anime) error {
	var AnimesTable = ksql.NewTable("anime")

	var animeResult domain.Anime
	err := s.db.QueryOne(ctx, &animeResult, "SELECT * FROM anime WHERE name = $1", anime.Name)
	if err != nil && err != ksql.ErrRecordNotFound {
		return domain.NewDatabaseErr("anime-service-create-query-anime-from-name", map[string]interface{}{})
	}

	if animeResult.Name == anime.Name {
		return domain.NewBadRequestErr("anime-service-create-duplicate-entry", map[string]interface{}{
			"error": "There is already an anime with this name registred",
		})
	}

	err = s.db.Insert(ctx, AnimesTable, &anime)
	if err != nil {
		return domain.NewDatabaseErr("anime-service-create-inserting-anime", map[string]interface{}{})
	}

	return nil
}

func (s Service) Update(ctx context.Context, anime domain.Anime) error {
	return nil
}
