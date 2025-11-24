package main

import (
	"context"
	"log"
	"os"
	"todo-api/internal/handler"
	"todo-api/internal/repository"
	"todo-api/internal/server"
	"todo-api/internal/service"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Initialization config error: %v\n", err)
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("env loading error: %v\n", err)
	}

	db, err := repository.NewPostgresDB(repository.DBConfig{
		User: viper.GetString("db.user"),
		Pass: os.Getenv("DB_PASS"),
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Name: viper.GetString("db.name"),
	})
	if err != nil {
		log.Fatalf("Unable to connect to db: %v\n", err)
	}
	defer db.Close(context.Background())

	log.Println("Successfully connected to database")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	router := handler.InitRouter(handlers)
	port := viper.GetString("port")
	srv := new(server.Server)

	if err := srv.Run(port, router); err != nil {
		log.Fatalf("Unable to start server: %v\n", err)
	} else {
		log.Printf("Server successfully started on port: %s\n", port)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
