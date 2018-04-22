package main

import (
	"github.com/krrrr38/go-freeehr/freeehr"
	"golang.org/x/oauth2"

	"fmt"
	"log"
	"os"
)

func main() {
	clientID := os.Getenv("FREEE_CLIENT_ID")
	if clientID == "" {
		log.Fatal("FREEE_CLIENT_ID env variable not found")
	}
	clientSecret := os.Getenv("FREEE_CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("FREEE_CLIENT_SECRET env variable not found")
	}
	conf := freeehr.Conf(clientID, clientSecret, "urn:ietf:wg:oauth:2.0:oob")

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\nEnter code: ", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("code: %v\n", code)

	token, err := conf.Exchange(oauth2.NoContext, code) // get access token and so on
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("access token: %v\n", token)

	client, _ := freeehr.NewClient(conf.Client(oauth2.NoContext, token))

	// get access token user information
	userResponse, resp, err := client.Users.GetMe()

	fmt.Printf("resp: %v\n", resp)
	if err != nil {
		fmt.Printf("got error: %v\n", err)
	} else {
		fmt.Printf("got user: %v\n", userResponse)
	}
}
