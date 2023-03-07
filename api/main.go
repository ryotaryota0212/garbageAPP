package main

import "api/routes"

func init() {
	// mysql接続
	// batch.GetExcelAndInsertCollection()
}

func main() {
	// apiサーバー起動
	routes.Server()
	// バッチ処理
	// batch.GetExcelAndInsert()
	// batch.GetExcelAndInsertCollection()
}
