package main

import (
	"context"
	"github.com/krrrr38/go-freeehr/freeehr"
	"golang.org/x/oauth2"

	"fmt"
	"os"
)

func main() {
	accessToken := os.Getenv("FREEE_ACCESS_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)

	client, _ := freeehr.NewClient(tc)

	// get access token user information
	userResponse, resp, err := client.Users.GetMe(ctx)

	fmt.Printf("resp: %v\n", resp)
	if err != nil {
		fmt.Printf("got error: %v\n", err)
	} else {
		fmt.Printf("got user: %v\n", userResponse)
	}
}
