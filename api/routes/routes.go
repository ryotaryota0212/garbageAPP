package routes

import (
	"api/firebase"

	"github.com/gin-gonic/gin"
)

func Server() {
	route := gin.Default()
	route.GET("/", Top)
	route.GET("/user", User)
	route.Run()
}
func Top(cont *gin.Context) {
	cont.JSON(200, gin.H{
		"message": "!!!!",
	})
}
func User(cont *gin.Context) {
	firebase.FirebaseIni()
	cont.JSON(200, gin.H{
		"message": "user",
	})
}
