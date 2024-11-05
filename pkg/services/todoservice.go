package services

import (
	"go-todo/config"
	db "go-todo/pkg/db/connector"
	"net/http"
	"sync"

	m "go-todo/pkg/db/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoSerice struct {
	DBInstance *gorm.DB
	Config     *config.Config
}

var OnceTodoService sync.Once

var SingleTodoServiceInstance *TodoSerice

type CreateTodoInput struct {
	Title           string `json:"t" binding:"required"`
	CompletedStatus bool   `json:"c"`
}

type UpdateTodoInput struct {
	Id              uint   `json:"i"`
	Title           string `json:"t"`
	CompletedStatus bool   `json:"c"`
}

func GetTodoService(config *config.Config) *TodoSerice {
	if SingleTodoServiceInstance == nil {
		OnceTodoService.Do(func() {
			dbconnector := db.GetDBConnectorInstance(config)
			SingleTodoServiceInstance = &TodoSerice{
				DBInstance: dbconnector.DB,
				Config:     config,
			}
		})
	}
	return SingleTodoServiceInstance
}

func (ts *TodoSerice) GetAllTodos(context *gin.Context) {
	var todoList []m.Todo
	result := ts.DBInstance.Find(&todoList).Order("id desc")
	if result.Error != nil {
		context.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"data": todoList})
}

func (ts *TodoSerice) CreateTodo(context *gin.Context) {
	var todoInput CreateTodoInput
	if err := context.ShouldBindJSON(&todoInput); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	todo := &m.Todo{Title: todoInput.Title, CompletedStatus: todoInput.CompletedStatus}
	result := ts.DBInstance.Create(todo)
	if result.Error != nil {
		context.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}

func (ts *TodoSerice) GetTodo(context *gin.Context) {
	id := context.Param("id")
	var todo m.Todo
	result := ts.DBInstance.First(&todo, id)
	if result.Error != nil {
		context.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"data": todo})
}

func (ts *TodoSerice) UpdateTodo(context *gin.Context) {
	var todoInput UpdateTodoInput
	if err := context.ShouldBindJSON(&todoInput); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var todo m.Todo
	if er := ts.DBInstance.Where("id = ?", todoInput.Id).First(&todo).Error; er != nil {
		context.AbortWithError(http.StatusBadRequest, er)
		return
	}
	result := ts.DBInstance.Model(&todo).Where("id = ?", todoInput.Id).Updates(todoInput)
	if result.Error != nil {
		context.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "TODO Updated successful", "data": todo})
}

func (ts *TodoSerice) DeleteTodo(context *gin.Context) {
	id := context.Param("id")
	var todo m.Todo
	if er := ts.DBInstance.Where("id = ?", id).First(&todo).Error; er != nil {
		context.AbortWithError(http.StatusBadRequest, er)
		return
	}
	result := ts.DBInstance.Delete(&todo, id)
	if result.Error != nil {
		context.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "TODO Deleted successful"})
}
