package repository

import (
	"context"
	"go-boilerplate/model"
	"go-boilerplate/pkg/constant"
)

func (r *Repository) CreateTodo(ctx context.Context, todo *model.Todo) error {
	err := r.db.WithContext(ctx).Table(constant.TodoTable).Create(&todo).Error
	return err
}
