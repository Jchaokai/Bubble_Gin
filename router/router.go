package router

import (
	"Bubble_Gin/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine{
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.RootHandler)
	taskGroup := r.Group("v1")
	{
		taskGroup.POST("/todo", controller.CreateTask)
		taskGroup.GET("/todo", controller.GetAllTask)
		taskGroup.GET("/todo/:id", controller.GetOneTask)
		taskGroup.PUT("/todo/:id", controller.UpdateTask)
		taskGroup.DELETE("/todo/:id", controller.DeleteOne)
	}
	return r
}