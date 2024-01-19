package repository

import (
	"context"
	"go-boilerplate/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}
type IRepository interface {
	CreateTodo(ctx context.Context, todo *model.Todo) error
}

func New(db *gorm.DB) IRepository {
	return &Repository{db: db}
}
