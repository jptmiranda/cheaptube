package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func CreateConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return conn, nil
}
