package routes

import (
	"api/controllers"
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func Server() {
	route := gin.Default()
	route.GET("/", Top)
	route.GET("/user", User)
	route.GET("/login/:id", Login)
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
	id := cont.Param("id")
	ctx := context.Background()
	opt := option.WithCredentialsFile("../public/auth/firebase-auth.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println(err)
		return
	}
	client, e := app.Auth(ctx)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(client)
	user := controllers.GetUser(ctx, client, id)
	fmt.Println(user)
	cont.JSON(200, gin.H{
		"user": user.UserInfo,
		"app": app,
	})
}

