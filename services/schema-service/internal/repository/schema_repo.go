package repository

import (
	"context"
	"schema-service/internal/db"
	"schema-service/internal/model"
)

func CreateSchema(s model.Schema) error {
	query := `
		INSERT INTO schemas (version, definition)
		VALUES ($1, $2)
	`

	_, err := db.Conn.Exec(context.Background(), query, s.Version, s.Definition)
	return err
}