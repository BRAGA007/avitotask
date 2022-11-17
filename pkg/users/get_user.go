package users

import (
	"avitotask/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetBalanceRequest struct {
	ID int `json:"id"`
}

func (h handler) GetUser(c *gin.Context) {
	body := GetBalanceRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, "Ошибка заполнения JSON")
		return
	}

	var user models.User

	if result := h.DB.First(&user, body.ID); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		c.JSON(http.StatusBadRequest, "Пользователя с данным ID не существует")
		return
	}

	c.JSON(http.StatusOK, &user)
}
