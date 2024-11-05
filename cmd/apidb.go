package cmd

import (

	ts "go-todo/pkg/services"
	// "net/http"
	// "strconv"
	"go-todo/config"

	"github.com/gin-gonic/gin"
)
var todoServ *ts.TodoSerice
func SetupDBRoutes(config *config.Config, router *gin.Engine) {
	todoServ = ts.GetTodoService(config)
	router.GET("/todos", todoServ.GetAllTodos)
	router.POST("/todos", todoServ.CreateTodo)
	router.GET("/todo/:id", todoServ.GetTodo)
	router.PATCH("/todos", todoServ.UpdateTodo)
	router.DELETE("/todo/:id", todoServ.DeleteTodo)
}

// func GetAllTodos(context *gin.Context) {
// 	todoList := todoSer.GetTodoList()
// 	context.IndentedJSON(http.StatusOK, todoList)
// }

// func CreateTodo(context *gin.Context) {
// 	var todo todo.Todo
// 	if err := context.BindJSON(&todo); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}
// 	er := todoSer.CreateTodo(todo)
// 	if er != nil {
// 		context.JSON(http.StatusContinue, gin.H{"message": er.Error()})
// 		return
// 	}
// 	context.IndentedJSON(http.StatusCreated, todo)
// }

// func GetTodo(context *gin.Context) {
// 	id , er := strconv.Atoi(context.Param("id"))
// 	if er != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": er.Error()})
// 		return
// 	}
// 	td, err := todoSer.GetTodo(int(id))
// 	if err != nil {
// 		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
// 		return 
// 	}
// 	context.IndentedJSON(http.StatusOK, td)
// }

// func UpdateTodo(context *gin.Context) {
// 	var todo todo.Todo
// 	if err := context.BindJSON(&todo); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}
// 	er := todoSer.UpdateTodo(todo)
// 	if er != nil {
// 		context.JSON(http.StatusNotFound, gin.H{"message": er.Error()})
// 		return
// 	}
// 	context.IndentedJSON(http.StatusOK, todo)
// }

// func DeleteTodo(context *gin.Context) {
// 	id , er := strconv.Atoi(context.Param("id"))
// 	if er != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": er.Error()})
// 		return
// 	}
// 	err := todoSer.DeleteTodo(id)
// 	if err != nil {
// 		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
// 		return
// 	}
// 	context.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted Successful"})
// }


