package routes

import (
	"api/controllers"
	"context"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func Server() {
	route := gin.Default()
	route.GET("/", Top)
	route.GET("/user", User)
	route.GET("/login", Login)
	route.GET("/prefectures", controllers.Prefectures)
	route.GET("/prefectures/:id", controllers.Municipality)
	route.GET("/garbage/:id", controllers.FetchGarbage)
	route.Run()
}
func Top(cont *gin.Context) {
	cont.JSON(200, gin.H{
		"message": "!!!!!!!!!!!!!!!!!!!!!!!!!!!!",
	})
}
func User(cont *gin.Context) {
	cont.JSON(200, gin.H{
		"message": "user",
	})
}
func Login(cont *gin.Context) {
	opt := option.WithCredentialsFile("../public/auth/firebase-auth.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return
	}
	cont.JSON(200, gin.H{
		"app": app,
	})
}

