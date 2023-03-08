package controllers_test

import (
	"api/database"
	"testing"
)

func TestFetchGarbage(t *testing.T) {
	db := database.ConnectDB()

	
	_,err := db.Query("select * from collections where local_number = 232017" )
	if err != nil {
		t.Errorf("not select collection by %v", err)
    }
}