package users

import (
	"avitotask/pkg/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DepositBalanceRequestBody struct {
	Id      uint `json:"user_id"`
	Deposit int  `json:"deposit"`
}

func (h handler) DepositBalance(c *gin.Context) {

	body := DepositBalanceRequestBody{}

	// получаем тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	var transaction models.Transaction

	if result := h.DB.First(&user, body.Id); result.Error != nil {
		user.Balance = body.Deposit
		user.ID = body.Id
		transaction.User_id = body.Id
		transaction.Desciption = "Пополнение баланса на сумму: " + strconv.Itoa(body.Deposit) + " копеек"
		fmt.Println(body.Deposit)

	} else {
		user.Balance += body.Deposit
		transaction.User_id = body.Id
		transaction.Desciption = "Пополнение баланса на сумму: " + strconv.Itoa(body.Deposit) + " копеек"
		fmt.Println(body.Deposit)
	}
	h.DB.Save(&user)
	h.DB.Save(&transaction)
	c.JSON(http.StatusOK, &user)
}
