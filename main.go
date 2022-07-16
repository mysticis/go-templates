package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/mysticis/golang-templates-app/api"
	db "github.com/mysticis/golang-templates-app/db/sqlc"
	"github.com/mysticis/golang-templates-app/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
