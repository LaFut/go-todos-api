package repository

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"todos-rest/logger"
	"todos-rest/model/entity"
)

type TodoListQueryBuilder struct {
	ListQueryBuilderInterface
	*BaseListQueryBuilder
}

func NewTodoListQueryBuilder(db *gorm.DB, logger *logger.Logger) ListQueryBuilderInterface {
	base := NewBaseListQueryBuilder(db, logger).(*BaseListQueryBuilder)
	return &TodoListQueryBuilder{BaseListQueryBuilder: base}
}

func (c TodoListQueryBuilder) ListQuery(parameters ListParametersInterface) *gorm.DB {
	query := c.BaseListQueryBuilder.ListQuery(parameters)

	val := reflect.ValueOf(parameters).Elem()
	if val.Kind() != reflect.Struct {
		c.logger.Fatal("Unexpected type of parameters for ListQuery")
		return query
	}

	field := val.FieldByName("parent_id")
	if field.IsValid() && field.Kind() == reflect.Int {
		query = query.Where("parent_id = ?", field.Int())
	}

	field = val.FieldByName("name")
	if field.IsValid() && field.Kind() == reflect.String {
		query = query.Where("name like ?", field.String()+"%")
	}

	return query
}

type TodoRepositoryInterface interface {
	CrudRepositoryInterface
	Children(id uint) (entity.InterfaceEntity, error)
}

type TodoParametersInterface interface {
	ListParametersInterface
}

type TodoParameters struct {
	*CrudListParameters
	Name     string
	ParentID int
}

type TodoRepository struct {
	*CrudRepository
	model entity.Todo
}

func (c TodoRepository) Children(id uint) (entity.InterfaceEntity, error) {
	items := reflect.New(reflect.SliceOf(reflect.TypeOf(c.GetModel()).Elem())).Interface()
	query := c.listQueryBuilder.ListQuery(TodoParameters{ParentID: int(id)})
	err := query.Find(items).Error
	return items, err
}

func NewTodoRepository(db *gorm.DB, logger *logger.Logger) TodoRepositoryInterface {
	var model entity.Todo
	queryBuilder := NewTodoListQueryBuilder(db, logger)
	repo := NewCrudRepository(db, &model, queryBuilder, logger).(*CrudRepository)
	return &TodoRepository{repo, model}
}
