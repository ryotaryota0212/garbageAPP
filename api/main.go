package main

import (
	"fmt"

	"api/database"
	"api/routes"
)

func init() {
	// mysql接続
	db := database.ConnectDB()
	fmt.Println("mysql成功", &db)
}

func main() {
	// apiサーバー起動
	routes.Server()
}
