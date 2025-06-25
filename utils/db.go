package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Conn, error) {
	godotenv.Load()
	
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("when creates new pool", err)
	}

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		fmt.Println("failed to acquire connection", err)
	}

	fmt.Println("connected to database")
	return conn, err
}