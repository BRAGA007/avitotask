package users

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Statement struct {
	ServiceId int
	Total     int
}
type GetRevenueStatementRequestBody struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

func (h handler) GetRevenueStatement(c *gin.Context) {
	body := GetRevenueStatementRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, "Ошибка заполнения JSON")
		return
	}

	var statements []Statement
	startofmonth := strconv.Itoa(body.Year) + "-" + strconv.Itoa(body.Month) + "-1 00:00:00"
	endofmonth := strconv.Itoa(body.Year) + "-" + strconv.Itoa(body.Month) + "-31 23:59:59"
	if result := h.DB.Table("revenues").Select("service_id, sum(amount) as total").Where("created_at >=? and created_at <= ?", startofmonth, endofmonth).Group("service_id").Scan(&statements); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		c.JSON(http.StatusBadRequest, "Данные за введенный период не найдены")
		return
	}
	csvFile, err := os.Create("pkg/revenues/" + strconv.Itoa(body.Year) + "-" + strconv.Itoa(body.Month) + ".csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	for _, s := range statements {
		var csvRow []string
		csvRow = append(csvRow, "Название услуги: "+strconv.Itoa(s.ServiceId)+";Общая сумма выручки за отчетный период "+strconv.Itoa(s.Total)+" копеек")
		writer.Write(csvRow)

	}
	c.JSON(http.StatusOK, "pkg/revenues/"+strconv.Itoa(body.Year)+"-"+strconv.Itoa(body.Month)+".csv")

	c.JSON(http.StatusOK, &statements)

}
