package database

import "github.com/jinzhu/gorm"

func Database() *gorm.DB {
	//open a database connection
	db, err := gorm.Open("mysql", "DB_USER:DB_PASSWORD@tcp(127.0.0.1:DB_PORT=3306)/sys?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
