package tests

import (
	"testing"
	"todos-rest/app"
	"todos-rest/model/entity"
	"todos-rest/model/repository"
	"todos-rest/model/service"
)

type TodoRepositoryMock struct {
	repository.CrudRepositoryInterface
	model entity.InterfaceEntity // Dynamic typing
}

func NewUserRepositoryMock() repository.TodoRepositoryInterface {
	var model entity.Todo
	repo := &TodoRepositoryMock{model: model}
	return repo
}

func (c TodoRepositoryMock) Find(id uint) (entity.InterfaceEntity, error) {
	item := &entity.Todo{TodoFields: &entity.TodoFields{Name: "test todo"}}
	item.ID = 2
	return item, nil
}

func (c TodoRepositoryMock) Children(id uint) (entity.InterfaceEntity, error) {
	var children = []entity.Todo{
		{
			TodoFields: &entity.TodoFields{
				Name:     "Child1",
				ParentId: 2,
			},
		},
		{
			TodoFields: &entity.TodoFields{
				Name:     "Child2",
				ParentId: 2,
			},
		},
	}
	return children, nil

}

func TestGetItem(t *testing.T) {
	a := app.App{}

	a.TodoRepository = NewUserRepositoryMock()
	a.TodoService = service.NewTodoService(a.TodoRepository, nil)
	todoService := a.TodoService

	item, err := todoService.GetItem(2)

	if err != nil {
		t.Error(err)
	}

	testItem := item.(*entity.Todo)
	if testItem.ID != 2 {
		t.Error("Id is not equals to 2")
	}

	if testItem.Name != "test todo" {
		t.Error("'test todo' expected")
	}
}
