package database

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	c := mysql.Config{
		DBName:    "garbage",
		User:      "admin",
		Passwd:    "admin",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}
	db, _ := sql.Open("mysql", c.FormatDSN())
	return db
}
