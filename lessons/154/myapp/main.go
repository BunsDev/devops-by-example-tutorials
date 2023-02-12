package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	_ "github.com/go-sql-driver/mysql"
)

var (
	maxClients = flag.Int("maxClients", 100, "Maximum number of virtual clients")
)

type metrics struct {
	duration *prometheus.SummaryVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		duration: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace:  "tester",
			Name:       "duration_seconds",
			Help:       "Duration of the request.",
			Objectives: map[float64]float64{0.9: 0.01, 0.99: 0.001},
		}, []string{"db", "operation"}),
	}
	reg.MustRegister(m.duration)
	return m
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Parse the command line into the defined flags
	flag.Parse()

	// Create Prometheus registry
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	// Create Prometheus HTTP server to expose metrics
	pMux := http.NewServeMux()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	pMux.Handle("/metrics", promHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":8081", pMux))
	}()

	dbUrl := "postgres://myapp:devops123@192.168.50.222:5432/benchmarks"
	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbpool.Close()

	db, err := sql.Open("mysql", "myappv2:devops123@tcp(192.168.50.87:3306)/benchmarks")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// See "Important settings" section.
	// db.SetConnMaxLifetime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	for i := 0; i < *maxClients; i++ {
		firstName, lastName := genName()

		err = insertAuthorToPostgres(dbpool, m, firstName, lastName)
		if err != nil {
			log.Fatalf("insertAuthor failed: %v", err)
		}

		err = insertAuthorToMysql(db, m, firstName, lastName)
		if err != nil {
			log.Fatalf("insertAuthor failed: %v", err)
		}

		time.Sleep(10 * time.Millisecond)
	}

	select {}
}

func insertAuthorToPostgres(p *pgxpool.Pool, m *metrics, firstName string, lastName string) error {
	now := time.Now()

	_, err := p.Exec(context.Background(), "INSERT INTO authors(first_name,last_name) VALUES($1,$2)", firstName, lastName)
	m.duration.With(prometheus.Labels{"db": "PostgreSQL", "operation": "write"}).Observe(time.Since(now).Seconds())

	return err
}

func insertAuthorToMysql(db *sql.DB, m *metrics, firstName string, lastName string) error {
	now := time.Now()

	_, err := db.Exec("INSERT INTO authors(first_name,last_name) VALUES(?,?)", firstName, lastName)
	m.duration.With(prometheus.Labels{"db": "MySQL", "operation": "write"}).Observe(time.Since(now).Seconds())

	return err
}

func genName() (string, string) {
	caser := cases.Title(language.English)

	firstName := caser.String(petname.Generate(1, ""))
	lastName := caser.String(petname.Generate(1, ""))

	return firstName, lastName
}
