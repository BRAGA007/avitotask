package users

import (
	"avitotask/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransferRequestBody struct {
	IdFrom   int `json:"id_from"`
	IdTo     int `json:"id_to"`
	Transfer int `json:"transfer"`
}

// Transfer godoc
// @Summary      Transferring money
// @Description  Transferring money between two selected users
// @Tags         Balance Interaction
// @Accept       json
// @Produce      json
// @Success      200  "Перевод выполнен"
// @Failure      400  "Ошибка заполнения JSON"
// @Router       /transfer [post]
func (h handler) Transfer(c *gin.Context) {

	body := TransferRequestBody{}

	// получаем тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var userto models.User
	var userfrom models.User
	var transactionto models.Transaction
	var transactionfrom models.Transaction

	if resultto := h.DB.First(&userto, body.IdTo); resultto.Error != nil {
		c.AbortWithError(http.StatusBadRequest, resultto.Error)
		c.JSON(http.StatusBadRequest, "Ошибка ввода данных")
		return

	}
	if resultfrom := h.DB.First(&userfrom, body.IdFrom); resultfrom.Error != nil {
		c.AbortWithError(http.StatusBadRequest, resultfrom.Error)
		c.JSON(http.StatusBadRequest, "Ошибка ввода данных")
		return

	}
	if body.IdTo == body.IdFrom {
		c.JSON(http.StatusBadRequest, "Нельзя переводить деньги самому себе")
		return
	}
	if userfrom.Balance < body.Transfer {
		c.JSON(http.StatusBadRequest, "Сумма перевода превышает баланс пользователя отправителя")
		return
	}

	if body.Transfer < +0 {
		c.JSON(http.StatusBadRequest, "Сумма перевода не может быть меньши либо равна нулю")
		return
	}
	userto.Balance += body.Transfer
	userfrom.Balance -= body.Transfer
	transactionfrom.Amount = body.Transfer
	transactionfrom.UserId = body.IdFrom
	transactionfrom.Description = "Отправлено: " + strconv.Itoa(body.Transfer) + " копеек " + "пользователю с ID " + strconv.Itoa(body.IdTo)
	transactionto.Amount = body.Transfer
	transactionto.UserId = body.IdTo
	transactionto.Description = "Получение: " + strconv.Itoa(body.Transfer) + " копеек " + "от пользователя с ID " + strconv.Itoa(body.IdFrom)
	h.DB.Save(&userto)
	h.DB.Save(&userfrom)
	c.JSON(http.StatusOK, "Перевод выполнен")

	h.DB.Save(&transactionto)
	h.DB.Save(&transactionfrom)
}
