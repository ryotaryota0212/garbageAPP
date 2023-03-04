package controllers

import (
	"api/database"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Garbage struct {
	Name string
	Division string
	Memo string
	SearchIndex string
	AddSearchIndex string
}

func FetchGarbage(cont *gin.Context) {
	id := cont.Param("id")
	db := database.ConnectDB()
    fmt.Println("mysql成功", &db)
	rows,err := db.Query("select name,division,memo,search_index,add_search_index from collections where local_number = " + id )
	if err != nil {
        fmt.Println(err)
        return
    }
	var garbages []Garbage
	for rows.Next() {
		var name,division,searchIndex string
		var memo,addSearchIndex sql.NullString
		if err := rows.Scan(&name,&division,&memo,&searchIndex,&addSearchIndex); err != nil {
			log.Fatal(err)
        }
		garbage := Garbage{
			Name: name,
			Division: division,
			Memo: memo.String,
			SearchIndex: searchIndex,
			AddSearchIndex: addSearchIndex.String,
		}
		fmt.Println(garbage)
		// arrayにpush
		garbages = append(garbages,garbage)
	}
	var count int
	e := db.QueryRow("select count(*) from collections where local_number = " + id ).Scan(&count)
	if e != nil {
		log.Fatal(err)
	}
	cont.JSON(200, gin.H{
		"data": garbages,
		"count": count,
	})
}