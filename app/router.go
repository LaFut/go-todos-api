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
	r.router.GET("/api/todos/:id", app.todoController.Get)
	r.router.GET("/api/todos", app.todoController.List)
	r.router.POST("/api/todos", app.todoController.Create)
	r.router.PUT("/api/todos/:id", app.todoController.Update)
	r.router.DELETE("/api/todos/:id", app.todoController.Delete)
}

func (r *Router) Run() {
	_ = r.router.Run(":" + viper.GetString("SERVER_PORT"))
}
