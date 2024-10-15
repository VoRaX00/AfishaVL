package main

import (
	"Afisha/internal/handler"
	"Afisha/internal/handler/extension"
	"Afisha/internal/infrastructure"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
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

	repos := extension.NewRepositories(db)
	services := extension.NewServices(repos)
	handlers := handler.NewHandler(services)

	server := new(Server)

	go func() {
		if err = server.Start(viper.GetString("server.port"), handlers.InitRouter()); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Info("Server started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
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
		Password: os.Getenv("db.password"),
		Database: viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
}
