package users

import (
	"avitotask/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetBalanceBodyRequest struct {
	ID int `json:"user_id"`
}

func (h handler) GetUser(c *gin.Context) {
	body := GetBalanceBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, "Ошибка заполнения JSON")
		return
	}

	var user models.User

	if result := h.DB.First(&user, body.ID); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		c.JSON(http.StatusBadRequest, "У пользователя с данным ID отсутствует баланс")
		return
	}

	c.JSON(http.StatusOK, &user)
}
