package freeehr

import (
	"golang.org/x/oauth2"
)

var endpoint = oauth2.Endpoint{
	AuthURL:  "https://secure.freee.co.jp/oauth/authorize",
	TokenURL: "https://api.freee.co.jp/oauth/token",
}

// Conf freee oauth2 conf
// If you try offline, you can get auth url by conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
// you can get access token by conf.Exchange(oauth2.NoContext, code) // get access token and so on
func Conf(clientID, clientSecret, redirectURL string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     endpoint,
	}
}
