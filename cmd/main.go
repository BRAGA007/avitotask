package main

import (
	"avitotask/pkg/common/models"
	"avitotask/pkg/users"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sys?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {

	db := Database()
	db.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Reservation{}, &models.Revenue{})
	router := gin.Default()

	users.RegisterRoutes(router, db)
	router.Run()

}
