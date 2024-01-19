package usecase

import (
	"go-boilerplate/repository"
	"go-boilerplate/usecase/todo"
)

type UseCase struct {
	TodoUseCase todo.IUseCase
}

func New(repo repository.IRepository) *UseCase {
	return &UseCase{
		TodoUseCase: todo.NewTodoUseCase(repo),
	}
}
