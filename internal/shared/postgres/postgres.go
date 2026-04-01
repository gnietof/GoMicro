package postgres

import (
	"context"
	"fmt"
	"os"

	_ "github.com/ibmdb/go_ibm_db"
	"github.com/jackc/pgx/v5"
)

type PostgresClient struct {
	Conn *pgx.Conn
}

func NewPostgresClient(ctx context.Context) (*PostgresClient, error) {

	connStr := buildConnString()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	// defer conn.Close(ctx)

	return &PostgresClient{Conn: conn}, nil
}

func buildConnString() string {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	pwd := os.Getenv("POSTGRES_PWD")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/postgres?sslmode=disable",
		user, pwd, host)

	return connStr
}
