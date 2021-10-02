package anime

import (
	"encoding/json"

	"github.com/brenoassp/animelist-go/domain"
	"github.com/valyala/fasthttp"
)

type Controller struct {
	animeService domain.AnimeService
}

func NewController(animeService domain.AnimeService) Controller {
	return Controller{
		animeService: animeService,
	}
}

func (c Controller) Create(ctx *fasthttp.RequestCtx) {
	var input struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
		NumEpisodes *int    `json:"num_episodes"`
	}

	err := json.Unmarshal(ctx.PostBody(), &input)
	if err != nil {
		// return json with validation error
	}

}
