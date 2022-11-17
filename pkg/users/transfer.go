package users

import (
	"avitotask/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransferRequestBody struct {
	Id_From  int `json:"id_from"`
	Id_To    int `json:"id_to"`
	Transfer int `json:"transfer"`
}

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

	if resultto := h.DB.First(&userto, body.Id_To); resultto.Error != nil {
		c.AbortWithError(http.StatusBadRequest, resultto.Error)
		c.JSON(http.StatusBadRequest, "Ошибка ввода данных")
		return

	}
	if resultfrom := h.DB.First(&userfrom, body.Id_From); resultfrom.Error != nil {
		c.AbortWithError(http.StatusBadRequest, resultfrom.Error)
		c.JSON(http.StatusBadRequest, "Ошибка ввода данных")
		return

	}
	if body.Id_To != body.Id_From {
		if userfrom.Balance >= body.Transfer {
			if body.Transfer > 0 {
				userto.Balance += body.Transfer
				userfrom.Balance -= body.Transfer
				transactionfrom.User_id = body.Id_From
				transactionfrom.Description = "Отправлено: " + strconv.Itoa(body.Transfer) + " копеек " + "пользователю с ID " + strconv.Itoa(int(body.Id_To))
				transactionto.User_id = body.Id_To
				transactionto.Description = "Получение: " + strconv.Itoa(body.Transfer) + " копеек " + "от пользователя с ID " + strconv.Itoa(int(body.Id_From))
				h.DB.Save(&userto)
				h.DB.Save(&userfrom)
				c.JSON(http.StatusOK, &userto)
				c.JSON(http.StatusOK, &userfrom)

				h.DB.Save(&transactionto)
				h.DB.Save(&transactionfrom)
			} else {
				c.JSON(http.StatusBadRequest, "Сумма перевода не может быть меньши либо равна нулю")
			}

		} else {
			c.JSON(http.StatusBadRequest, "Сумма перевода превышает баланс пользователя отправителя")
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, "Нельзя переводить деньги самому себя")
		return
	}
}
