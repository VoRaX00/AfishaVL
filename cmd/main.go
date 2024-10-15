package main

import (
	"Afisha/internal/infrastructure"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	cfg := initDBConfig()

	db, err := infrastructure.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		_ = db.Close()
	}()

	if err = goose.Up(db.DB, "./storage/migrations"); err != nil {
		logrus.Fatal(err)
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initDBConfig() infrastructure.Config {
	return infrastructure.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		Database: viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
}
