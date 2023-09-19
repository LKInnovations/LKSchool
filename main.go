package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка переменных окружения из файла .env
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файлов: %v\n", err)
	}
	fmt.Println()

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Ошибка подключения дб: %v\n", err)
	}
	defer dbPool.Close()
}
