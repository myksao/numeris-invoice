package db

import (
	"embed"
	"fmt"
	"invoice/config"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewDB(c *config.Config, fs embed.FS) (*sqlx.DB, error) {

	source, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s search_path=numeris",
		c.Database.Host,
		c.Database.Port,
		c.Database.Username,
		c.Database.Name,
		c.Database.Password,
	)

	// Open a connection to create the schema if it doesn't exist
	schemadb, err := sqlx.Connect(c.Database.Driver, dsn)
	if err != nil {
		return nil, err
	}
	defer schemadb.Close()

	// Create schema if not exists
	_, err = schemadb.Exec("CREATE SCHEMA IF NOT EXISTS numeris")
	if err != nil {
		return nil, err
	}
	/** Start Migrate the database */
	driver, _ := postgres.WithInstance(sqlx.MustOpen(c.Database.Driver, dsn).DB, &postgres.Config{
		MigrationsTable:       fmt.Sprintf("\"%s\".\"%s\"", c.Database.Schema, "migrations"),
		MigrationsTableQuoted: true,
	})
	migrate, err := migrate.NewWithInstance("iofs", source, "postgresql", driver)
	if err != nil {
		log.Fatal("migrate: ", err)
	}
	if err := migrate.Up(); err != nil {
		log.Println("migrate: ", err)
	}
	/** End of migration */

	db, err := sqlx.Connect(c.Database.Driver, dsn)
	if err != nil {
		return nil, err
	}

	// Set the search path globally
    _, err = db.Exec("SET search_path TO numeris")
    if err != nil {
        return nil, err
    }

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
