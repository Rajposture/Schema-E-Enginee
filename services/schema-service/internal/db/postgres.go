package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() error {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/asee")
	if err != nil {
		return err
	}

	Conn = conn
	fmt.Println("Connected to DB")
	return nil
}