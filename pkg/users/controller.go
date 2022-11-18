package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/api/v1")
	{
		routes.POST("/statement", h.GetRevenueStatement)
		routes.POST("/", h.GetUser)
		routes.POST("/history", h.GetHistoryBalance)
		routes.POST("/withdraw", h.WithDrawBalance)
		routes.POST("/deposit", h.DepositBalance)
		routes.POST("/transfer", h.Transfer)
		routes.POST("/reserve", h.ReserveBalanceAndRevenueRecognition)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
