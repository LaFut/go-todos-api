package repository

import (
	"github.com/jinzhu/gorm"
	"reflect"
	"todos-rest/logger"
	"todos-rest/model/entity"
)

type ListParametersInterface interface{}

type CrudListParameters struct {
}

type ListQueryBuilderInterface interface {
	ListQuery(parameters ListParametersInterface) *gorm.DB
}

type BaseListQueryBuilder struct {
	db     *gorm.DB
	logger *logger.Logger
	ListQueryBuilderInterface
}

func NewBaseListQueryBuilder(db *gorm.DB, logger *logger.Logger) ListQueryBuilderInterface {
	return &BaseListQueryBuilder{db: db, logger: logger}
}

func (c BaseListQueryBuilder) ListQuery(parameters ListParametersInterface) *gorm.DB {
	query := c.db

	val := reflect.ValueOf(parameters).Elem()
	if val.Kind() != reflect.Struct {
		c.logger.Fatal("Unexpected type of parameters for ListQuery")
		return query
	}
	//common page limits here if needed
	return query
}

type CrudRepositoryInterface interface {
	BaseRepositoryInterface
	GetModel() entity.InterfaceEntity
	Find(id uint) (entity.InterfaceEntity, error)
	List(parameters ListParametersInterface) (entity.InterfaceEntity, error)
	Create(item entity.InterfaceEntity) entity.InterfaceEntity
	Update(item entity.InterfaceEntity) entity.InterfaceEntity
	Delete(id uint) error
}

type CrudRepository struct {
	CrudRepositoryInterface
	*BaseRepository
	model            entity.InterfaceEntity // Dynamic typing
	listQueryBuilder ListQueryBuilderInterface
}

func NewCrudRepository(db *gorm.DB, model entity.InterfaceEntity, listQueryBuilder ListQueryBuilderInterface, logger *logger.Logger) CrudRepositoryInterface {
	repo := NewBaseRepository(db, logger).(*BaseRepository)
	return &CrudRepository{
		BaseRepository:   repo,
		model:            model,
		listQueryBuilder: listQueryBuilder,
	}
}

func (c CrudRepository) GetModel() entity.InterfaceEntity {
	return c.model
}

func (c CrudRepository) Find(id uint) (entity.InterfaceEntity, error) {
	item := reflect.New(reflect.TypeOf(c.GetModel()).Elem()).Interface()
	err := c.db.First(item, id).Error
	return item, err
}

func (c CrudRepository) List(parameters ListParametersInterface) (entity.InterfaceEntity, error) {
	items := reflect.New(reflect.SliceOf(reflect.TypeOf(c.GetModel()).Elem())).Interface()
	query := c.listQueryBuilder.ListQuery(parameters)
	err := query.Find(items).Error
	return items, err
}

func (c CrudRepository) Create(item entity.InterfaceEntity) entity.InterfaceEntity {
	c.db.Create(item)
	return item
}

func (c CrudRepository) Update(item entity.InterfaceEntity) entity.InterfaceEntity {
	c.db.Save(item)
	return item
}

func (c CrudRepository) Delete(id uint) error {
	item, err := c.Find(id)
	if err != nil {
		return err
	}
	c.db.Delete(item)
	return nil
}
