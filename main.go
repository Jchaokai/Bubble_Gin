package main

import (
	"Bubble_Gin/dao"
	"Bubble_Gin/entity"
	"Bubble_Gin/router"
)

func main() {
	if e := dao.Conn("mysql");e!=nil{
		panic(e)
	}
	defer dao.Close()
	entity.InitTable()
	r := router.SetUpRouter()
	_ = r.Run(":8080")
}
