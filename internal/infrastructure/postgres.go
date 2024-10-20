package infrastructure

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Host,
			config.Port, config.Username, config.Database, config.Password, config.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	return db, err
}
