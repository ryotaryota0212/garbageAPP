package controllers

import (
	"api/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Prefecture struct {
	Id int
	Name string
	Division string
}

func Prefectures(cont *gin.Context) {
	db := database.ConnectDB()
    fmt.Println("mysql成功", &db)
	rows,err := db.Query("select * from global_code")
	if err != nil {
        fmt.Println(err)
        return
    }
	var prefectures []Prefecture
	for rows.Next() {
		var id int
		var name,division string
		if err := rows.Scan(&id,&name,&division); err != nil {
			log.Fatal(err)
        }
		prefecture := Prefecture{
			Id: id,
			Name: name,
			Division: division,
		}
		fmt.Println(prefecture)
		// arrayにpush
		prefectures = append(prefectures,prefecture)
	}
	cont.JSON(200, prefectures)
}