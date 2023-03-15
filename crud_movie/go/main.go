package main

import (
	"database/sql"
	"log"

	api "github.com/sam-val/languages/crud_movie/go/api"
	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
	"github.com/sam-val/languages/crud_movie/go/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can not load config ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to DB ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server ", err)
	}

}