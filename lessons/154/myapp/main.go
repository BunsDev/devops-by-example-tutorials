package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	dbUrl := "postgres://myapp:devops123@192.168.50.222:5432/benchmarks"
	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	for i := 1; i < 5; i++ {
		firstName, lastName := genName()

		err = insertAuthor(dbpool, firstName, lastName)
		if err != nil {
			log.Fatalf("insertAuthor failed: %v", err)
		}
	}

}

func insertAuthor(p *pgxpool.Pool, firstName string, lastName string) error {
	_, err := p.Exec(context.Background(), "INSERT INTO authors(first_name,last_name) values($1,$2)", firstName, lastName)
	return err
}

func genName() (string, string) {
	caser := cases.Title(language.English)

	firstName := caser.String(petname.Generate(1, ""))
	lastName := caser.String(petname.Generate(1, ""))

	return firstName, lastName
}
