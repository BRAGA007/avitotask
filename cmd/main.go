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

func Database(DBURL string) *gorm.DB {
	fmt.Println(DBURL)
	db, err := gorm.Open("mysql", DBURL)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Reservation{}, &models.Revenue{})
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
	DB_HOST := viper.Get("DB_HOST").(string)
	//DB_URL := DB_USER + ":" + DB_PASSWORD + "@tcp(127.0.0.1:" + DB_PORT + ")/" + DB_NAME
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db := Database(DBURL)
	db.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Reservation{}, &models.Revenue{})
	router := gin.Default()

	users.RegisterRoutes(router, db)
	router.Run()

}
