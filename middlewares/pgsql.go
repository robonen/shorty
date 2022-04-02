package middlewares

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

const databaseUrl = "postgres://shorty:secretd@localhost:5432/shorty"

type (
	Handler struct {
		DB *pgxpool.Pool
	}
)

const (
	Database = "secret"
	User     = "user"
	Password = "pass"
)

func Test() {
	// this returns connection pool
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// to close DB pool
	defer dbPool.Close()
}
