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
	item, err := c.repository.Find(id)

	if err != nil {
		return item, err
	}

	val := reflect.ValueOf(item).Elem()
	if val.Kind() != reflect.Struct {
		c.logger.Fatal("")
		return item, err
	}

	children, err := c.repository.Children(id)
	reflect.ValueOf(item).Elem().FieldByName("Children").Set(reflect.ValueOf(children).Elem())

	return item, err
}
