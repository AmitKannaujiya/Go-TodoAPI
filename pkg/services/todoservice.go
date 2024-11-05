package services

import (
	"go-todo/config"
	db "go-todo/pkg/db/connector"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoSerice struct {
	DBInstance *gorm.DB
	Config *config.Config
}

var OnceTodoService sync.Once

var SingleTodoServiceInstance *TodoSerice

func GetTodoService(config *config.Config) *TodoSerice {
	if SingleTodoServiceInstance == nil {
		OnceTodoService.Do(func() {
			dbconnector := db.GetDBConnectorInstance(config)
			SingleTodoServiceInstance = &TodoSerice{
				DBInstance: dbconnector.DB,
				Config: config,
			}
		})
	}
	return SingleTodoServiceInstance
}

func (ts *TodoSerice) GetAllTodos(contex *gin.Context) {

}

func (ts *TodoSerice) CreateTodo(contex *gin.Context) {
	
}

func (ts *TodoSerice) GetTodo(contex *gin.Context) {
	
}

func (ts *TodoSerice) UpdateTodo(contex *gin.Context) {
	
}

func (ts *TodoSerice) DeleteTodo(contex *gin.Context) {
	
}