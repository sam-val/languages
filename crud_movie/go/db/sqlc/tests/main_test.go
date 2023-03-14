package tests 

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	db "github.com/sam-val/languages/crud_movie/go/db/sqlc"
	"github.com/sam-val/languages/crud_movie/go/util"
)

var testQueries *db.Queries

func TestMain(t *testing.M) {
	conf, err := util.LoadConfig("./../../..")	
	if err != nil {
		log.Fatal("Could not load configs ", err.Error())
	}

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("Could not connect to DB ", err.Error())
	}

	testQueries = db.New(conn)
	os.Exit(t.Run())

}