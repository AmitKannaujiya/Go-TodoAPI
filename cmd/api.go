package cmd

import (
	todo "go-todo/models"
	todoService "go-todo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
var todoSer *todoService.TodoService
func SetupRoutes(router *gin.Engine) {
	todoSer = todoService.GetTodoService()
	router.GET("/todos", GetAllTodos)
	router.POST("/todos", CreateTodo)
	router.GET("/todo/:id", GetTodo)
	router.PATCH("/todos", UpdateTodo)
	router.DELETE("/todo/:id", DeleteTodo)
}

func GetAllTodos(context *gin.Context) {
	todoList := todoSer.GetTodoList()
	context.IndentedJSON(http.StatusOK, todoList)
}

func CreateTodo(context *gin.Context) {
	var todo todo.Todo
	if err := context.BindJSON(&todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	er := todoSer.CreateTodo(todo)
	if er != nil {
		context.JSON(http.StatusContinue, gin.H{"message": er})
		return
	}
	context.IndentedJSON(http.StatusCreated, todo)
}

func GetTodo(context *gin.Context) {
	id , er := strconv.Atoi(context.Param("id"))
	if er != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": er})
		return
	}
	td, err := todoSer.GetTodo(int(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err})
		return 
	}
	context.IndentedJSON(http.StatusOK, td)
}

func UpdateTodo(context *gin.Context) {
	var todo todo.Todo
	if err := context.BindJSON(&todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	er := todoSer.UpdateTodo(todo)
	if er != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": er})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func DeleteTodo(context *gin.Context) {
	id , er := strconv.Atoi(context.Param("id"))
	if er != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": er})
		return
	}
	err := todoSer.DeleteTodo(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted Successful"})
}


