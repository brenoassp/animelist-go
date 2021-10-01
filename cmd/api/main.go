package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brenoassp/animelist-go/infra/env"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/vingarcia/ksql"
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

	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		fmt.Printf("Hello, world! Requested path is %q", ctx.Path())
	})

	r.NotFound = func(ctx *fasthttp.RequestCtx) {
		response := struct {
			Message string
		}{
			Message: "Invalid route",
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			fmt.Println("Error marshalling route not found response")
			os.Exit(1)
		}

		ctx.SetStatusCode(http.StatusNotFound)
		ctx.SetContentType("application/json")
		ctx.SetBody(responseJSON)
	}

	if err := fasthttp.ListenAndServe("127.0.0.1:8090", r.Handler); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
