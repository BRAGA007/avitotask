package main

import (
	"avitotask/docs"
	"avitotask/pkg/common/models"
	"avitotask/pkg/users"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

func Database(DB_URL string) *gorm.DB {
	//open a database connection
	db, err := gorm.Open("mysql", DB_URL)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// @title       AvitoTask
// @version     1.0
// @host        localhost:8080
// @BasePath    /api/v1
// @description Microservice for working with user balance

func main() {
	docs.SwaggerInfo.Title = "AvitoTask"
	docs.SwaggerInfo.Description = "Microservice for working with user balance"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	DB_USER := viper.Get("DB_USER").(string)
	DB_NAME := viper.Get("DB_NAME").(string)
	DB_PASSWORD := viper.Get("DB_PASSWORD").(string)
	DB_PORT := viper.Get("DB_PORT").(string)
	DB_URL := DB_USER + ":" + DB_PASSWORD + "@tcp(127.0.0.1:" + DB_PORT + ")/" + DB_NAME
	fmt.Println(DB_URL)
	db := Database(DB_URL)
	db.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Reservation{}, &models.Revenue{})
	router := gin.Default()

	users.RegisterRoutes(router, db)
	router.Run()

}
