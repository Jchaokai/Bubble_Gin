package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	DB *gorm.DB
)


func Conn(DBSource string)  (err error){
	connString := newConnString("root", "123",  "localhost", 3306, "db_test")
	DB, err = gorm.Open(DBSource, connString)
	return
}
func Close(){
	DB.Close()
}

func newConnString(user string,pwd string,ip string,port int,db string) string{
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",user, pwd,ip,port,db)
}
func InitModel(values ...interface{}){
	DB.AutoMigrate(values)
}