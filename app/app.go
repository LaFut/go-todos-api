package app

import (
	"github.com/jinzhu/gorm"
	"sync"
	"todos-rest/controller"
	"todos-rest/logger"
	"todos-rest/model/db"
	"todos-rest/model/repository"
	"todos-rest/model/service"
)

type App struct {
	logger         *logger.Logger
	db             *gorm.DB
	router         *Router
	todoService    service.TodoServiceInterface
	todoRepository repository.TodoRepositoryInterface
	todoController *controller.TodoController
	appOnce        sync.Once
}

var app App

func NewApp() *App {
	app := App{}
	return &app
}

func (a *App) Init() {
	a.logger = logger.NewLogger()
	a.db = db.NewDb()
	a.router = NewRouter()
	a.todoRepository = repository.NewTodoRepository(a.db, a.logger)
	a.todoService = service.NewTodoService(a.todoRepository, a.logger)
	a.todoController = controller.NewTodoController(a.todoService, a.logger)

	a.router.InitRouter(a)
}

func (a *App) Run() {
	a.router.Run()
}
