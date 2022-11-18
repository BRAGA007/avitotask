package users

import (
	"github.com/jinzhu/gorm"
	"math"
)

type Param struct {
	DB    *gorm.DB
	Page  int
	Limit int
}

type Paginator struct {
	History     interface{} `json:"history"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
	TotalRecord int         `json:"total_record"`
	TotalPage   int         `json:"total_page"`
}

// Paging godoc
// @Summary      Pagination
// @Description  Create pagination
// @Tags         Pagination
// @Accept       json
// @Produce      json
// @Success      200  {object}  Paginator
// @Failure      400  "Ошибка заполнения JSON"
func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int
	var offset int

	go countRecords(db, result, done, &count)

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	db.Limit(p.Limit).Offset(offset).Find(result)
	<-done

	paginator.TotalRecord = count
	paginator.History = result
	paginator.Page = p.Page

	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}
