# go-freeehr

:pray: Under Development, Please Do Not Use :pray:

go api client for [freee HR](https://www.freee.co.jp/hr/).

- [API Specification]((https://www.freee.co.jp/hr/api).

## Usage

- `FREEE_ACCESS_TOKEN=${your_access_token} go run example/user/main.go`
  - if yo do not have access token, see [document](https://support.freee.co.jp/hc/ja/articles/115000145263-freee-API%E3%81%AE%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E3%83%88%E3%83%BC%E3%82%AF%E3%83%B3%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B).

```go
package main

import (
	"context"
	"path/to/go-freeehr/freeehr"
	"golang.org/x/oauth2"

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
}
```

## LICENSE

This library is distributed under the BSD-style license. Almost all base code are written in [google/go-github](https://github.com/google/go-github).
