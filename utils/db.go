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

	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	database := os.Getenv("PGDATABASE")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)

	pool, err := pgxpool.New(context.Background(), connectionString)
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