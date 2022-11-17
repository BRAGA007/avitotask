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
	routes.POST("/res", h.GetRevenueStatement)
	routes.POST("/", h.GetUser)
	routes.POST("/history", h.GetHistoryBalance)
	routes.POST("/withdraw", h.WithDrawBalance)
	routes.POST("/deposit", h.DepositBalance)
	routes.POST("/transfer", h.Transfer)
	routes.POST("/reserve", h.ReserveBalanceAndRevenueRecognition)
}
