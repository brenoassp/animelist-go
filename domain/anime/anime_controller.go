package anime

import (
	"encoding/json"
	"fmt"

	"github.com/brenoassp/animelist-go/domain"
	atreugo "github.com/savsgio/atreugo/v11"
)

type Controller struct {
	animeService domain.AnimeService
}

func NewController(animeService domain.AnimeService) Controller {
	return Controller{
		animeService: animeService,
	}
}

func (c Controller) Create(ctx *atreugo.RequestCtx) error {
	var input struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
		NumEpisodes *int    `json:"num_episodes"`
	}

	err := json.Unmarshal(ctx.PostBody(), &input)
	if err != nil {
		// return json with validation error
		ctx.SetUserValue("error", fmt.Errorf("error unmarshaling body from request to create new anime"))
	}
	return nil
}
