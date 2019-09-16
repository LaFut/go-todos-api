package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"todos-rest/logger"
	"todos-rest/model/repository"
	"todos-rest/model/service"
)

type ParametersReceiverInterface interface {
	Receive(context *gin.Context) (repository.ListParametersInterface, error)
}

type BaseParametersReceiver struct {
	logger *logger.Logger
	ParametersReceiverInterface
}

func NewBaseParametersReceiver(logger *logger.Logger) ParametersReceiverInterface {
	return &BaseParametersReceiver{logger: logger}
}

func (c BaseParametersReceiver) Receive(context *gin.Context) (repository.ListParametersInterface, error) {
	parameters := &repository.CrudListParameters{}

	return parameters, nil
}

type CrudControllerInterface interface {
	BaseControllerInterface
	Create(context *gin.Context)
	Get(context *gin.Context)
	List(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type CrudController struct {
	CrudControllerInterface
	*BaseController
	service            service.CrudServiceInterface
	parametersReceiver ParametersReceiverInterface
}

func NewCrudController(service service.CrudServiceInterface, parametersReceiver ParametersReceiverInterface, logger *logger.Logger) *CrudController {
	controller := NewBaseController(logger)
	return &CrudController{BaseController: controller, service: service, parametersReceiver: parametersReceiver}
}

func (c CrudController) Get(context *gin.Context) {
	recordId, err := strconv.Atoi(context.Params.ByName("id"))
	if err != nil {
		c.response(context, gin.H{"Entity": nil, "Status": "error"}, http.StatusBadRequest)
		return
	}

	entity, err := c.service.GetItem(uint(recordId))

	if err != nil {
		c.response(context, gin.H{"Entity": nil, "Status": "error"}, http.StatusNotFound)
		return
	}

	c.response(context, gin.H{"Entity": entity, "Status": "ok"}, http.StatusOK)
}

func (c CrudController) List(context *gin.Context) {
	parameters, err := c.parametersReceiver.Receive(context)

	if err != nil {
		c.response(context, gin.H{"Entities": nil, "Status": "error"}, http.StatusBadRequest)
	}

	entities, err := c.service.GetList(parameters)

	if err != nil {
		c.response(context, gin.H{"Entities": nil, "Status": "error"}, http.StatusNotFound)
		return
	}

	c.response(context, gin.H{"Entities": entities, "Status": "ok"}, http.StatusOK)
}

func (c CrudController) Create(context *gin.Context) {
	model := c.service.GetModel()
	item := reflect.New(reflect.TypeOf(model).Elem()).Interface()
	if err := context.ShouldBindJSON(item); err != nil {
		c.response(context, gin.H{"Entity": nil, "Status": "error"}, http.StatusBadRequest)
		return
	}
	item = c.service.Create(item)
	c.response(context, gin.H{"Entity": item, "Status": "ok"}, http.StatusOK)
}

func (c CrudController) Update(context *gin.Context) {
	recordId, err := strconv.Atoi(context.Params.ByName("id"))
	if err != nil {
		c.response(context, gin.H{"Entity": nil, "Status": "error"}, http.StatusBadRequest)
		return
	}

	entity, err := c.service.GetItem(uint(recordId))
	if err != nil {
		c.response(context, gin.H{"Entity": nil, "Status": "error"}, http.StatusNotFound)
		return
	}

	if err := context.ShouldBindJSON(entity); err != nil {
		c.response(context, gin.H{"Entity": nil, "Status": "error"}, http.StatusBadRequest)
		return
	}
	entity = c.service.Update(entity)

	c.response(context, gin.H{"Entity": entity, "Status": "ok"}, http.StatusOK)
}

func (c CrudController) Delete(context *gin.Context) {
	recordId, err := strconv.Atoi(context.Params.ByName("id"))
	if err != nil {
		c.response(context, gin.H{"Entity": nil, "Status": "error"}, http.StatusBadRequest)
		return
	}

	err = c.service.Delete(uint(recordId))
	if err != nil {
		c.response(context, gin.H{"Entity": nil, "Status": "error"}, http.StatusNotFound)
		return
	}

	c.response(context, gin.H{"Entity": nil, "Status": "ok"}, http.StatusOK)
}
