package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todos-rest/logger"
	"todos-rest/model/repository"
	"todos-rest/model/service"
)

type TodoListParametersReceiver struct {
	*BaseParametersReceiver
}

func NewTodoListParametersReceiver(logger *logger.Logger) ParametersReceiverInterface {
	base := NewBaseParametersReceiver(logger).(*BaseParametersReceiver)
	return &TodoListParametersReceiver{BaseParametersReceiver: base}
}

func (c TodoListParametersReceiver) Receive(context *gin.Context) (repository.ListParametersInterface, error) {
	crudParams, _ := c.BaseParametersReceiver.Receive(context)

	parameters := &repository.TodoParameters{
		CrudListParameters: crudParams.(*repository.CrudListParameters),
	}
	parentId, err := strconv.Atoi(context.Params.ByName("parentid"))
	if err == nil {
		parameters.ParentID = parentId
	}

	if err := context.ShouldBindQuery(parameters); err != nil {
		return crudParams, err
	}

	return parameters, nil
}

type TodoController struct {
	*CrudController
	service service.TodoServiceInterface
}

func NewTodoController(service service.TodoServiceInterface, logger *logger.Logger) *TodoController {
	parametersReceiver := NewTodoListParametersReceiver(logger)
	controller := NewCrudController(service, parametersReceiver, logger)
	return &TodoController{CrudController: controller, service: service}
}

func (c TodoController) Get(context *gin.Context) {
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
