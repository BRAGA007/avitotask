package db

import "github.com/jinzhu/gorm"

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sys?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
