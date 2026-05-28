package service

import (
	"schema-service/internal/model"
	"schema-service/internal/repository"
)

func CreateSchema(s model.Schema) error {
	return repository.CreateSchema(s)
}