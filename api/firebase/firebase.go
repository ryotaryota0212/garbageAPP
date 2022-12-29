package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func FirebaseIni() {
	opt := option.WithCredentialsFile("./auth_private.json")
	fmt.Println(opt)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(app, "成功！！")
	return
}
