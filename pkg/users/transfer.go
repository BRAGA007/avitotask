package users

import (
	"avitotask/pkg/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TransferRequestBody struct {
	Id_From  uint `json:"id_from"`
	Id_To    uint `json:"id_to"`
	Transfer int  `json:"transfer"`
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
		c.JSON(http.StatusBadRequest, "Пользователя получателся с таким ID не существует")
		return

	}
	if resultfrom := h.DB.First(&userfrom, body.Id_From); resultfrom.Error != nil {
		c.AbortWithError(http.StatusBadRequest, resultfrom.Error)
		c.JSON(http.StatusBadRequest, "Пользователя отправителя с таким ID не существует")
		return

	}
	if userfrom.Balance >= body.Transfer {
		userto.Balance += body.Transfer
		userfrom.Balance -= body.Transfer
		transactionfrom.User_id = body.Id_From
		transactionfrom.Desciption = "Отправлено: " + strconv.Itoa(body.Transfer) + " копеек " + "пользователю с ID " + strconv.Itoa(int(body.Id_To))
		transactionto.User_id = body.Id_To
		transactionto.Desciption = "Получение: " + strconv.Itoa(body.Transfer) + " копеек " + "от пользователя с ID " + strconv.Itoa(int(body.Id_From))
		h.DB.Save(&userto)
		h.DB.Save(&userfrom)
		c.JSON(http.StatusOK, &userto)
		c.JSON(http.StatusOK, &userfrom)

		h.DB.Save(&transactionto)
		h.DB.Save(&transactionfrom)
	} else {
		c.JSON(http.StatusBadRequest, "Сумма перевода превышает баланс пользователя отправителя")
	}
}
