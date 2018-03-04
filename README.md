# go-freeehr

[![Build Status](https://travis-ci.org/krrrr38/go-freeehr.svg?branch=master)](https://travis-ci.org/krrrr38/go-freeehr)

go api client for [freee HR](https://www.freee.co.jp/hr/).

- [API Specification](https://www.freee.co.jp/hr/api).

## Usage

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

## Example

- `FREEE_ACCESS_TOKEN=${your_access_token} go run example/user/main.go`
  - if yo do not have access token, see [document](https://support.freee.co.jp/hc/ja/articles/115000145263-freee-API%E3%81%AE%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E3%83%88%E3%83%BC%E3%82%AF%E3%83%B3%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B).

## Development

- install dependency

```sh
> make dep
```

- run test

```sh
> make test
```

## LICENSE

This library is distributed under the BSD-style license. Almost all code base are written in [google/go-github](https://github.com/google/go-github).
