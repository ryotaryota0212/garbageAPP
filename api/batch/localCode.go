package batch

import (
	"api/database"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

type LocalCode struct {
    GlobalCode int `json:"global_code"`
	LocalCode string `json:"local_code"`
	Name string `json:"name"`
}

func GetExcelAndInsert () {
	f,err := excelize.OpenFile("../public/documents/localCode.xlsx")
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

    rows, err := f.GetRows(sheetFirst)
    if err != nil {
        fmt.Println(err)
        return
    }
    db := database.ConnectDB()
    fmt.Println("mysql成功", &db)
    for _, row := range rows[1:] {
        // for _, colCell := range row {
        //     fmt.Print(colCell, "\t")
        // }
		fmt.Println()
		fmt.Println(row[0][0:2])

        // global_code, _ := strconv.Atoi(row[0][0:2])
        // local_code := row[0]
        name := row[2]
		// fmt.Println(local_code,name)
        if len(name) != 0 {
            furigana := row[4]
            // ins, err := db.Prepare("INSERT INTO local_code(global_code,local_code,name) VALUES(?,?,?)")
            ins, err := db.Prepare("update local_code set furigana = ? where name = ?")
            if err != nil {
                log.Fatal(err)
            }
            // ins.Exec(global_code,local_code,name)
            ins.Exec(furigana,name)
        }
    }
}