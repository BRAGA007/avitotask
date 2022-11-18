package users

import (
	"avitotask/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DepositBalanceRequestBody struct {
	Id      int `json:"user_id"`
	Deposit int `json:"deposit"`
}

// DepositBalance godoc
// @Summary      Deposit Balance
// @Description  Deposit Balance
// @Tags         Balance Interaction
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Failure      400  "Ошибка заполнения JSON"
// @Router       /deposit [post]
func (h handler) DepositBalance(c *gin.Context) {

	body := DepositBalanceRequestBody{}

	// получаем тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	var transaction models.Transaction

	if body.Id <= 0 {
		c.JSON(http.StatusBadRequest, "Введите ID больше 0")
		return
	}
	if body.Deposit <= 0 {
		c.JSON(http.StatusBadRequest, "Сумма пополнения длжна быть больше 0")
		return
	}

	_ = h.DB.First(&user, body.Id)
	user.ID = body.Id
	user.Balance += body.Deposit
	transaction.UserId = body.Id
	transaction.Description = "Пополнение баланса на сумму: " + strconv.Itoa(body.Deposit) + " копеек"
	transaction.Amount = body.Deposit
	h.DB.Save(&user)
	h.DB.Save(&transaction)
	c.JSON(http.StatusOK, &user)
}
