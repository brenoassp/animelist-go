package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/brenoassp/animelist-go/domain"
	"github.com/brenoassp/animelist-go/domain/anime"
	"github.com/brenoassp/animelist-go/infra/env"
	"github.com/vingarcia/ksql"

	atreugo "github.com/savsgio/atreugo/v11"
)

func main() {
	type config struct {
		postgresURIWithDB          string
		postgresMaxOpenConnections int
	}

	conf := config{
		postgresURIWithDB:          env.MustGetString("POSTGRES_URI_WITH_DB"),
		postgresMaxOpenConnections: env.GetInt("POSTGRES_MAX_OPEN_CONNECTIONS", 5),
	}

	db, err := ksql.New("postgres", conf.postgresURIWithDB, ksql.Config{
		MaxOpenConns: conf.postgresMaxOpenConnections,
	})
	if err != nil {
		fmt.Printf(
			"Error creating ksql db for connectionString: %s using %d connections. Ksql error: %v",
			conf.postgresURIWithDB,
			conf.postgresMaxOpenConnections,
			err,
		)
		os.Exit(1)
	}

	animeService := anime.NewService(db)
	animeController := anime.NewController(animeService)

	server := atreugo.New(atreugo.Config{
		Addr: "127.0.0.1:8090",
	})
	server.UseAfter(handleErrors)

	server.GET("/", func(ctx *atreugo.RequestCtx) error {
		return ctx.TextResponse("Hello World")
	})

	animeGroup := server.NewGroupPath("/anime")
	animeGroup.POST("", func(ctx *atreugo.RequestCtx) error {
		return animeController.Create(ctx)
	})

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func handleErrors(ctx *atreugo.RequestCtx) error {
	ctxErr := ctx.UserValue("error")
	if ctxErr == nil {
		return ctx.Next()
	}

	domainErr, ok := ctxErr.(domain.DomainErr)
	if !ok {
		// gera log de erro
		genericError(ctx)
		return nil
	}

	httpStatus := http.StatusInternalServerError

	switch domainErr.Code {
	case "BadRequestErr":
		httpStatus = http.StatusBadRequest
	case "NotFoundErr":
		httpStatus = http.StatusNotFound
	}

	errData, err := json.Marshal(domainErr.Data)
	if err != nil {
		// gera log de erro
		genericError(ctx)
		return nil
	}

	ctx.SetStatusCode(httpStatus)
	ctx.SetBody(errData)
	return ctx.Next()
}

func genericError(ctx *atreugo.RequestCtx) {
	output := struct {
		Error string `json:"error"`
	}{
		Error: "internal error",
	}
	outputJSON, err := json.Marshal(output)
	if err != nil {
		// gera log de erro
	}
	ctx.SetStatusCode(http.StatusInternalServerError)
	ctx.SetBody(outputJSON)
}
