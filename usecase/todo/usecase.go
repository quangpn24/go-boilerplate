package todo

import (
	"context"
	"go-boilerplate/payload"
	"go-boilerplate/repository"
)

type UseCase struct {
	repo repository.IRepository
}
type IUseCase interface {
	CreateTodo(ctx context.Context, req payload.CreateTodoRequest) error
}

func NewTodoUseCase(repo repository.IRepository) IUseCase {
	return &UseCase{
		repo: repo,
	}
}
