package controllers

import (
	"context"
	"log"

	"firebase.google.com/go/auth"
)

func GetUser(ctx context.Context, client *auth.Client, uid string) (*auth.UserRecord) {
    u, err := client.GetUser(ctx, uid)
    if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
    }
    log.Printf("Successfully fetched user data: %#v\n", u.UserInfo)

    return u
}