package database

import (
	"log"
	"os"
	"time"

	"github.com/deeep8250/vibecheck-api/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DBinit() {

	//mention DSN
	dsn := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"

	// for docker commpose we trying looping several time because docker compose took several seconds to start postgres
	var err error
	for range 5 {
		config.PostgresDB, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("error while connecting with database", err.Error())
	}

	// finding the migrate
	m, err := migrate.New("file://db/migrations", dsn)
	if err != nil {
		log.Fatal("migration failed", err.Error())
	}

	//execute the founded migration
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("migration failed ", err.Error())
	}

	log.Println("migration successful")
	log.Println("database connection established")
}
