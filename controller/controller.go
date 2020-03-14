package controller

import (
	"Bubble_Gin/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func CreateTask(c *gin.Context) {
	var task entity.Task
	c.BindJSON(&task)
	if e := entity.Create(&task); e != nil {
		c.JSON(http.StatusOK, gin.H{"error": e.Error(),})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": task,})
	}
}

func GetAllTask(c *gin.Context) {
	lists, err := entity.FindAll()
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK, lists)
	}
}
func GetOneTask(c *gin.Context) {
	id := c.Param("id")
	if task, err := entity.GetOne(id);err !=nil{
		c.JSON(http.StatusOK,gin.H{"error":err})
	}else{
		c.JSON(http.StatusOK, task)
	}
}
func UpdateTask(c *gin.Context) {
	id, e := c.Params.Get("id")
	if !e {
		c.JSON(http.StatusOK, gin.H{"error": "id not found"})
		return
	}
	task, err := entity.GetOne(id)
	if  err != nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}
	c.BindJSON(&task)
	if e := entity.UpdateOne(task);e!=nil{
		c.JSON(http.StatusOK,gin.H{"error":e.Error()})
	}else{
		c.JSON(http.StatusOK, task)
	}
}

func DeleteOne(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error":"id is not found"})
		return
	}
	if e := entity.Delete(id);e!=nil{
		c.JSON(http.StatusOK,gin.H{"error":e.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id:"deleted"})
	}
}