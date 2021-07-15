package main

import (
	"context"
	"os"
	"time"

	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var app = cli.App("pgconn_check", "postgres connection checker")
var (
	appPostgresConnection = app.String(cli.StringOpt{
		Name:   "postgresConnection",
		Desc:   "Postgres connection string",
		EnvVar: "POSTGRES_CONNECTION",
		Value:  "postgres://postgres:example@localhost:5432/db?sslmode=disable",
	})
)

func startApplication() {
	log.Printf("testing %v\n", *appPostgresConnection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := pgxpool.Connect(ctx, *appPostgresConnection)
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func main() {
	app.Action = startApplication
	if err := app.Run(os.Args); err != nil {
		log.Fatalln("[ERR]", err)
	}
}
