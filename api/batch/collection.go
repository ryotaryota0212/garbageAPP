package batch

import (
	"api/database"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

type garbageCollection struct {
	Division string `json:"division"`
	Name string `json:"name"`
	Memo string `json:"memo"`
	Search_index string `json:"search_index"`
	Add_search_index string `json:"add_search_index"`
	Local_number string `json:"local_number"`
}

func GetExcelAndInsertCollection () {
	f,err := excelize.OpenFile("../public/documents/豊橋市分別早見表.xlsx")
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
	sheetFirst := f.GetSheetName(0)
	fmt.Println(sheetFirst)

    rows, err := f.GetRows(sheetFirst)
    if err != nil {
        fmt.Println(err)
        return
    }
    db := database.ConnectDB()
    fmt.Println("mysql成功", &db)
	// defer db.Close()
    for _, row := range rows[2:] {
        // for _, colCell := range row {
        //     fmt.Print(colCell, "\t")
        // }
		// fmt.Println()
		fmt.Println(row)
        name := row[1]
		// division := row[2]
		memo := ""
		if row[3] != "null" {
			memo = row[3]
		}
		addIndex := ""
		if row[4] != "null" {
			addIndex = row[4]
		}
		// if len(row[3]) != 0  {
		// }
        // index := row[0]
		// addIndex := ""
		// if len(row[4]) != 0  {
		// 	addIndex = row[4]
		// }
		// ins, err := db.Prepare("INSERT INTO collections(name,division,search_index,local_number) VALUES(?,?,?,?)")
		ins, err := db.Prepare("UPDATE collections set memo = ?,add_search_index = ? where name = ?")
		if err != nil {
			log.Fatal(err)
		}
		// ins.Exec(name,division,index,232017)
		ins.Exec(memo,addIndex,name)
    }
}