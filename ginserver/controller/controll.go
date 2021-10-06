package controller

import (
	"ginserver/database"
	"ginserver/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTodo(ctx *gin.Context) {
	DB := database.DB
	todo := new(model.Todo)
	todo.Id = uuid.New()
	ctx.BindJSON(&todo)
	d := DB.Create(&todo)

	ctx.IndentedJSON(http.StatusOK, d)
}

func GetTodo(ctx *gin.Context) {
	DB := database.DB
	var todos []model.Todo
	DB.Find(&todos)

	ctx.IndentedJSON(http.StatusOK, todos)
}
