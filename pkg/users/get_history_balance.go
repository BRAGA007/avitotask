package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetHistoryBalanceBodyRequest struct {
	User_id  int    `json:"user_id"`
	Sort_by  string `json:"sort_by"`
	Order_by string `json:"order_by"`
}
type TransactionResponse struct {
	Created_at  string `json:"date"`
	Description string `json:"description"`
}

func (h handler) GetHistoryBalance(c *gin.Context) {
	body := GetHistoryBalanceBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, "Ошибка заполнения JSON")
		return
	}

	available_request_sort_body := []string{"created_at", "amount"}
	available_request_order_body := []string{"asc", "desc"}
	order_body := []string{body.Order_by}
	sort_body := []string{body.Sort_by}

	if CheckSortData(available_request_sort_body, sort_body) && CheckSortData(available_request_order_body, order_body) {
		order_by := body.Sort_by + " " + body.Order_by

		var transactionResponse []TransactionResponse
		if result := h.DB.Table("transactions").Select("description, DATE_FORMAT(created_at, '%Y-%m-%d %h:%m:%s') as created_at").Where("user_id = ?", body.User_id).Order(order_by).Scan(&transactionResponse); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			c.JSON(http.StatusBadRequest, "Данные за введенный период не найдены")
			return
		}
		c.JSON(http.StatusOK, "История баланса пользователя с ID: "+strconv.Itoa(body.User_id))
		c.JSON(http.StatusOK, &transactionResponse)
	} else {
		c.JSON(http.StatusBadRequest, "Данные сортировки введены неправильно")
	}
}
func CheckSortData(available_request_body, request_body []string) bool {
	flag := false
	for _, body := range request_body {
		flag = false
		for _, available := range available_request_body {
			fmt.Println(body, available)
			if body == available {
				flag = true
			}
		}
	}
	fmt.Println(flag)
	return flag
}
