package controllers

import (
	"api/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Municipal struct {
	LocalCode string
	Name string
}

func Municipality(cont *gin.Context) {
	id := cont.Param("id")
	db := database.ConnectDB()
    fmt.Println("mysql成功", &db)
	rows,err := db.Query("select local_code,name from local_code where global_code = " + id )
	if err != nil {
        fmt.Println(err)
        return
    }
	var municipals []Municipal
	for rows.Next() {
		var localCode,name string
		if err := rows.Scan(&localCode,&name); err != nil {
			log.Fatal(err)
        }
		municipal := Municipal{
			LocalCode: localCode,
			Name: name,
		}
		fmt.Println(municipal)
		// arrayにpush
		municipals = append(municipals,municipal)
	}
	var count int
	e := db.QueryRow("select count(*) from local_code where global_code = " + id ).Scan(&count)
	if e != nil {
		log.Fatal(err)
	}
	cont.JSON(200, gin.H{
		"data": municipals,
		"count": count,
	})
}