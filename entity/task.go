package entity

import "Bubble_Gin/dao"

type Task struct {
	ID int	`json:"id"`
	Title string	`json:"title"`
	Status bool	`json:"status"`
}

func (Task) TableName() string{
	return "bubble_tasks"
}

func InitTable(){
	dao.DB.AutoMigrate(&Task{})
}

//CRUD
func Create(task *Task)(err error){
	err = dao.DB.Create(&task).Error
	return
}
func GetOne(id string)(task *Task,err error){
	if err = dao.DB.Where("id=?",id).First(&task).Error;err!=nil{
		return nil,err
	}
	return
}
func UpdateOne(task *Task) error{
	return dao.DB.Save(task).Error
}

func FindAll() (lists []Task,err error){
	if err = dao.DB.Find(&lists).Error;err!=nil{
		return nil,err
	}
	return
}
func Delete(id string) error{
	return dao.DB.Where("id=?",id).Delete(&Task{}).Error
}
