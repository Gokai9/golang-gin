package router

import (
	"ginserver/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("api/create", controller.CreateTodo)
	r.GET("api/todos", controller.GetTodo)

	return r
}
