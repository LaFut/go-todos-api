// Application
//
// Application description
//
//     Schemes: http
//     Host: localhost:80
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "todos-rest/app/description"
)

type IRouterInterface interface {
	InitRouter() *gin.Engine
}

type Router struct {
	router *gin.Engine
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) InitRouter(app *App) {
	r.router = gin.Default()

	// Display Swagger documentation
	r.router.StaticFile("doc/swagger.json", "doc/swagger.json")
	config := &ginSwagger.Config{
		URL: "/doc/swagger.json", //The url pointing to API definition
	}
	// use ginSwagger middleware to
	r.router.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	// swagger:route GET /api/todos/:id todo getTodo
	//
	// Todo
	//
	// Get todo data
	//
	//     Responses:
	//       200: TodoResponse
	r.router.GET("/api/todos/:id", app.TodoController.Get)

	// swagger:route GET /api/todos todo GetTodos
	//
	// Todo list
	//
	// Get todo list data
	//
	//     Responses:
	//       200: TodoResponse
	r.router.GET("/api/todos", app.TodoController.List)

	// swagger:route POST /api/todos todo CreateTodo
	//
	// New todo
	//
	// Create new todo
	//
	//     Responses:
	//       200: TodoResponse
	r.router.POST("/api/todos", app.TodoController.Create)

	// swagger:route PUT /api/todos/:id todo UpdateTodo
	//
	// Update todo
	//
	// Update existing todo
	//
	//     Responses:
	//       200: TodoResponse
	r.router.PUT("/api/todos/:id", app.TodoController.Update)

	// swagger:route DELETE /api/todos/:id todo DeleteTodo
	//
	// Delete todo
	//
	// Delete existing todo
	//
	//     Responses:
	//       200:
	r.router.DELETE("/api/todos/:id", app.TodoController.Delete)
}

func (r *Router) Run() {
	_ = r.router.Run(":" + viper.GetString("SERVER_PORT"))
}
