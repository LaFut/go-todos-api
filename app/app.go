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
	TodoService    service.TodoServiceInterface
	TodoRepository repository.TodoRepositoryInterface
	TodoController *controller.TodoController
	appOnce        sync.Once
}

func NewApp() *App {
	app := App{}
	return &app
}

func (a *App) Init() {
	a.logger = logger.NewLogger()
	a.db = db.NewDb()
	a.router = NewRouter()
	a.TodoRepository = repository.NewTodoRepository(a.db, a.logger)
	a.TodoService = service.NewTodoService(a.TodoRepository, a.logger)
	a.TodoController = controller.NewTodoController(a.TodoService, a.logger)

	a.router.InitRouter(a)
}

func (a *App) Run() {
	a.router.Run()
}
