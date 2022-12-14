package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Host     string `required:"true"`
	Port     int    `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Dbname   string `envconfig:"DATABASE" required:"true"`

	MaxConnectionLifetime time.Duration `envconfig:"DB_MAX_CONN_LIFE_TIME" default:"300s"`
	MaxOpenConnection     int           `envconfig:"DB_MAX_OPEN_CONNECTION" default:"100"`
	MaxIdleConnection     int           `envconfig:"DB_MAX_IDLE_CONNECTION" default:"10"`
}

func (p Postgres) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)
}

func OpenDatabaseConnection(pg Postgres) (*sql.DB, error) {
	dbConn, err := sql.Open("postgres", pg.ConnectionString())
	if err != nil {
		return nil, err
	}

	dbConn.SetConnMaxLifetime(pg.MaxConnectionLifetime)
	dbConn.SetMaxOpenConns(pg.MaxOpenConnection)
	dbConn.SetMaxIdleConns(pg.MaxIdleConnection)

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
