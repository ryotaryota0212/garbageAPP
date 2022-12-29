package main

import (
	"fmt"

	"api/database"
	"api/firebase"
	"api/routes"
)

func init() {
	// mysql接続
	db := database.ConnectDB()
	fmt.Println("mysql成功", &db)
	// firebase接続
	firebase.FirebaseIni()
}

func main() {
	// apiサーバー起動
	routes.Server()
}
