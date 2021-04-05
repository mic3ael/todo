package service

import (
	"github.com/mic3ael/todo/pkg/model"
	"github.com/mic3ael/todo/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList)(int, error)
	GetAll(userId int)([]model.TodoList, error)
	GetById(userId, listId int)(model.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, id int, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item model.TodoItem)(int, error)
	GetAll(userId, listId int) ([]model.TodoItem, error)
	GetById(userId, itemId int) (model.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input model.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:	NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}