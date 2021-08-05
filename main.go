package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v37/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	user, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	// Rate.Limit should most likely be 5000 when authorized.
	log.Printf("Rate: %#v", resp.Rate)

	fmt.Printf("\n%v\n", github.Stringify(user))
}
