package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/balances")
	//routes.POST("/", h.AddUser)
	//routes.GET("/", h.GetUsers)
	routes.GET("/", h.GetUser)
	routes.POST("/", h.DepositBalance)
	routes.POST("/transfer", h.Transfer)
	routes.POST("/reserve", h.ReserveBalanceAndRevenueRecognition)
}
