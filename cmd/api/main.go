package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-api/internal/handler"
	"todo-api/internal/repository"
	"todo-api/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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

	var (
		user   = viper.GetString("db.user")
		pass   = os.Getenv("DB_PASS")
		host   = viper.GetString("db.host")
		dbport = viper.GetString("db.port")
		dbname = viper.GetString("db.name")
	)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, dbport, dbname)
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to db: %v\n", err)
	}
	defer db.Close(context.Background())

	log.Println("Successfully connected to database")

	todoRepo := repository.NewRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running!"))
	})
	r.Post("/todos", todoHandler.CreateTodo)

	port := ":" + viper.GetString("port")
	log.Printf("Server started successfully on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
