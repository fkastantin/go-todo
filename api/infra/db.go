package infra

import (
    "fmt"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func GetDBConnection() *gorm.DB{
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_NAME")

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, db_name)

	db, err := gorm.Open(mysql.Open(url))

	if err != nil {
		panic("Failed to connect to database!")
	}

	fmt.Println("Database connection is established")

	//sql, err := db.DB()
	//sql.SetMaxIdleConns()

	return db
}