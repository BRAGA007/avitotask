package users

import (
	"avitotask/pkg/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReserveBalanceAndRevenueRecognitionRequestBody struct {
	User_Id    int `json:"user_id"`
	Service_Id int `json:"service_id"`
	Order_Id   int `json:"order_id"`
	Cost       int `json:"cost"`
}

func (h handler) ReserveBalanceAndRevenueRecognition(c *gin.Context) {

	body := ReserveBalanceAndRevenueRecognitionRequestBody{}

	// получаем тело запроса
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(body.Order_Id)

	var user models.User
	var transaction models.Transaction
	var reservation models.Reservation
	var revenue models.Revenue
	if result := h.DB.First(&user, body.User_Id); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		c.JSON(http.StatusBadRequest, "У заказчика с таким ID отсутствует баланс")
		return
	} else {
		if result := h.DB.First(&reservation, body.Order_Id); result.Error != nil {
			if body.Cost > 0 && body.Service_Id > 0 && body.Order_Id > 0 {

				reservation.User_Id = body.User_Id
				reservation.Service_Id = body.Service_Id
				reservation.Order_Id = body.Order_Id
				reservation.Cost = body.Cost
				reservation.Status = "Заказ не подтвержден"
			} else {
				c.JSON(http.StatusBadRequest, "Сумма покупки, ID услуги, ID заказа не могут быть меньше либо равны нулю")
				return
			}

		} else {

			if reservation.User_Id == body.User_Id {
				if reservation.Status == "Заказ не подтвержден" {
					if reservation.Service_Id == body.Service_Id {
						if reservation.Cost == body.Cost {
							if user.Balance >= reservation.Cost {
								user.Balance -= reservation.Cost
								h.DB.Save(&user)
								reservation.Status = "Заказ подтвержден"
								transaction.User_id = reservation.User_Id
								transaction.Desciption = "Подтверждение заказа: " + strconv.Itoa(int(reservation.Order_Id)) + " на сумму " + strconv.Itoa(reservation.Cost) + " копеек"
								h.DB.Save(&transaction)
								revenue.User_Id = reservation.User_Id
								revenue.Service_Id = reservation.Service_Id
								revenue.Sum = reservation.Cost
								revenue.Order_Id = reservation.Order_Id
								h.DB.Save(&revenue)
							} else {

								c.JSON(http.StatusBadRequest, "Недостаточно средств на балнсе для подтверждения заказа")
								return
							}
						} else {
							c.JSON(http.StatusBadRequest, "Сумма заказа не совпадает с суммой резервации")
							return
						}
					} else {
						c.JSON(http.StatusBadRequest, "Усулга резервации не совпадает с услугой подтверждения")
						return
					}

				} else {
					c.JSON(http.StatusBadRequest, "Невозможно подтвердить уже оплаченный заказ")
					return
				}
			} else {
				c.JSON(http.StatusBadRequest, "Данный заказ принадлежит другому пользователю")
				return
			}
		}
		h.DB.Save(&reservation)

		c.JSON(http.StatusOK, &user)
	}

}
