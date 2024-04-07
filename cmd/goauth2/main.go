package main

import (
	"context"
	"log"

	"github.com/haton14/goauth2"
)

func main() {
	ctx := context.Background()
	oauth2Conf, err := goauth2.NewOAuth2Config()
	if err != nil {
		log.Fatal(err)
	}
	authorizationURL, err := goauth2.BuildAutorizationRequestURL(oauth2Conf)
	if err != nil {
		log.Fatal(err)
	}
	authzCode, err := goauth2.GetAuthorizationCode(ctx, authorizationURL, oauth2Conf.RedirectURL)
	if err != nil {
		log.Fatal(err)
	}
	accessToken, err := goauth2.GetAccessToken(ctx, oauth2Conf, authzCode)
	if err != nil {
		log.Fatal(err)
	}
	photos, err := goauth2.GetProtectedResource(ctx, oauth2Conf, accessToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(photos)
}
