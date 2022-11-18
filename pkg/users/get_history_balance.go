package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetHistoryBalanceBodyRequest struct {
	UserId  int    `json:"user_id"`
	SortBy  string `json:"sort_by"`
	OrderBy string `json:"order_by"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}
type TransactionResponse struct {
	CreatedAt   string `json:"date"`
	Description string `json:"description"`
}

// GetHistoryBalance godoc
// @Summary      History Balance
// @Description  Shows all balance history for selected user
// @Tags         Balance Interaction
// @Accept       json
// @Produce      json
// @Success      200  {object}   GetHistoryBalanceBodyRequest
// @Failure      400  "Ошибка заполнения JSON"
// @Failure      404  "Данные за введенный период не найдены"
// @Router       /history [post]
func (h handler) GetHistoryBalance(c *gin.Context) {
	body := GetHistoryBalanceBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, "Ошибка заполнения JSON")
		return
	}
	if body.UserId <= 0 {
		c.JSON(http.StatusBadRequest, "Введите ID пользователя больше нуля")
		return
	}

	available_request_sort_body := map[string]struct{}{"created_at": {}, "amount": {}}
	available_request_order_body := map[string]struct{}{"asc": {}, "desc": {}}

	_, isSortByValid := available_request_sort_body[body.SortBy]
	_, isOrderByValid := available_request_order_body[body.OrderBy]
	if (isSortByValid && (len(body.OrderBy) == 0 || isOrderByValid)) || len(body.SortBy+body.OrderBy) == 0 {
		order_by := body.SortBy + " " + body.OrderBy
		if len(body.SortBy) != 0 && len(body.OrderBy) == 0 {
			order_by = body.SortBy
		}
		if len(body.SortBy) == 0 && len(body.OrderBy) == 0 {
			order_by = ""
		}

		var transactionResponse []TransactionResponse
		if result := h.DB.Table("transactions").Select("description, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i') as created_at").Where("user_id = ?", body.UserId).Order(order_by).Scan(&transactionResponse); result.Error != nil {
			c.AbortWithError(http.StatusNotFound, result.Error)
			c.JSON(http.StatusBadRequest, "Данные за введенный период не найдены")
			return
		}
		if body.Limit == 0 {
			body.Limit = 5
		}

		paginator := Paging(&Param{
			DB:    h.DB.Table("transactions").Select("description, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i') as created_at").Where("User_id = ?", body.UserId).Order(order_by).Scan(&transactionResponse),
			Page:  body.Page,
			Limit: body.Limit,
		}, &transactionResponse)
		c.JSON(http.StatusOK, &paginator)
	} else {
		c.JSON(http.StatusBadRequest, "Данные сортировки введены неправильно")
	}
}
