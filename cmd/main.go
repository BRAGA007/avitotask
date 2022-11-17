package main

import (
	"avitotask/pkg/common/models"
	"avitotask/pkg/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database(DB_URL string) *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", DB_URL)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	DB_URL := viper.Get("DB_URL").(string)
	fmt.Println(DB_URL)
	db := Database(DB_URL)
	db.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Reservation{}, &models.Revenue{})
	router := gin.Default()

	users.RegisterRoutes(router, db)
	router.Run()

}
