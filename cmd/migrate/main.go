package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/shamank/ai-marketplace-stats-service/internal/config"
)

func main() {

	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "", "path to config file")

	var migrationsPath string
	flag.StringVar(&migrationsPath, "migrations", "", "path to migrations")

	flag.Parse()

	cfg, err := config.LoadConfig(cfgPath)

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	))
	if err != nil {
		panic(err)
	}

	flag.Parse()

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"postgres",
		driver,
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		panic(err)
	}
}
