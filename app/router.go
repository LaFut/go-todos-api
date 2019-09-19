package app

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
	r.router.GET("/api/todos/:id", app.TodoController.Get)
	r.router.GET("/api/todos", app.TodoController.List)
	r.router.POST("/api/todos", app.TodoController.Create)
	r.router.PUT("/api/todos/:id", app.TodoController.Update)
	r.router.DELETE("/api/todos/:id", app.TodoController.Delete)
}

func (r *Router) Run() {
	_ = r.router.Run(":" + viper.GetString("SERVER_PORT"))
}
