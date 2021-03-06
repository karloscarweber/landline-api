// forego run go run db/test-data.go

package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/asm-products/landline-api/models"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "[DATABASE] ", log.Lmicroseconds))

	dbmap.AddTableWithName(models.Team{}, "teams").SetKeys(true, "id")

	team := &models.Team{
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		Email:             "jake@ooo.com",
		EncryptedPassword: "s3kr3th4sh",
		SSOUrl:            "http://localhost:8989/sso",
		SSOSecret:         "41fe7589256fd058b3f56bc71a56ebad3b1d6b86e027a73a02db0e3a0524f9d4",
		Slug:              "test-dev",
	}

	err = dbmap.Insert(team)
	if err != nil {
		log.Fatal(err)
	}
}
