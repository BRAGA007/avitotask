package users

import (
	"avitotask/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WithDrawBalanceRequestBody struct {
	Id         int `json:"user_id"`
	Withdrawal int `json:"withdrawal"`
}

func (h handler) WithDrawBalance(c *gin.Context) {

	body := WithDrawBalanceRequestBody{}

	// получаем тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	var transaction models.Transaction

	if result := h.DB.First(&user, body.Id); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		c.JSON(http.StatusBadRequest, "Ошибка ввода данных")
		return

	} else {
		if user.Balance >= body.Withdrawal {

			user.Balance -= body.Withdrawal
			transaction.User_id = body.Id
			transaction.Description = "Вывод средств на сумму: " + strconv.Itoa(body.Withdrawal) + " копеек"
		} else {
			c.JSON(http.StatusBadRequest, "Недостаточно средств для вывода")
			return
		}

	}
	h.DB.Save(&user)
	h.DB.Save(&transaction)
	c.JSON(http.StatusOK, &user)
}
