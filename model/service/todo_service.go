package service

import (
	"reflect"
	"todos-rest/logger"
	"todos-rest/model/entity"
	"todos-rest/model/repository"
)

type TodoServiceInterface interface {
	CrudServiceInterface
}

type TodoService struct {
	*CrudService
	repository repository.TodoRepositoryInterface
}

func NewTodoService(repository repository.TodoRepositoryInterface, logger *logger.Logger) TodoServiceInterface {
	crudService := NewCrudService(repository, logger).(*CrudService)
	service := &TodoService{crudService, repository}
	return service
}

func (c TodoService) GetItem(id uint) (entity.InterfaceEntity, error) {
	err, todo := c.repository.Find(id)

	if err != nil {
		return err, todo
	}

	val := reflect.ValueOf(todo).Elem()
	if val.Kind() != reflect.Struct {
		c.logger.Fatal("")
		return err, todo
	}

	//val.Children = c.repository.Children(id)

	return err, todo
}
