package anime

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func (c Controller) Create(ctx *atreugo.RequestCtx) {
	var input struct {
		Name        *string `json:"name"`
		Description *string `json:"description"`
		NumEpisodes *int    `json:"num_episodes"`
	}

	err := json.Unmarshal(ctx.PostBody(), &input)
	if err != nil {
		// return json with validation error
		ctx.SetUserValue(
			"error",
			domain.NewBadRequestErr("anime-ctrl-create-invalid-body", map[string]interface{}{
				"error": "invalid body. Request body should be a json",
			}))
		return
	}

	if input.Name == nil {
		ctx.SetUserValue(
			"error",
			domain.NewBadRequestErr("anime-ctrl-create-invalid-body-missing-name", map[string]interface{}{
				"error": "Missing field `name` in body.",
			}))
		return
	}

	anime := domain.Anime{
		Name:        input.Name,
		Description: input.Description,
		NumEpisodes: input.NumEpisodes,
	}

	err = c.animeService.Create(ctx, anime)
	if err != nil {
		ctx.SetUserValue("error", err)
		return
	}

	ctx.SetStatusCode(http.StatusCreated)
	ctx.SetBody([]byte(``))
}

func (c Controller) Update(ctx *atreugo.RequestCtx) {
	animeIDStr, ok := ctx.UserValue("id").(string)
	if !ok {
		ctx.SetUserValue("error",
			domain.NewBadRequestErr(
				"anime-ctrl-update-error-converting-path-parameter",
				map[string]interface{}{
					"error": "invalid id given in path parameter",
				},
			))
		return
	}

	animeID, err := strconv.Atoi(animeIDStr)
	if err != nil {
		ctx.SetUserValue("error", domain.NewBadRequestErr(
			"anime-ctrl-update-invalid-anime-id-given",
			map[string]interface{}{
				"error": fmt.Sprintf("Anime id given in the path parameter should be an integer, but `%s` given", animeIDStr),
			},
		))
		return
	}

	var input struct {
		Name        *string `json:"name"`
		Description *string `json:"description"`
		NumEpisodes *int    `json:"num_episodes"`
	}

	err = json.Unmarshal(ctx.PostBody(), &input)
	if err != nil {
		// return json with validation error
		ctx.SetUserValue(
			"error",
			domain.NewBadRequestErr("anime-ctrl-update-invalid-body", map[string]interface{}{
				"error": "invalid body. Request body should be a json",
			}))
		return
	}

	anime := domain.Anime{
		ID:          animeID,
		Name:        input.Name,
		Description: input.Description,
		NumEpisodes: input.NumEpisodes,
	}

	err = c.animeService.Update(ctx, anime)
	if err != nil {
		ctx.SetUserValue("error", err)
		return
	}

	ctx.SetStatusCode(http.StatusAccepted)
	ctx.SetBody([]byte(``))
}
