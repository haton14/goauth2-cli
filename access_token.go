package goauth2

import (
	"context"

	"golang.org/x/oauth2"
)

func GetAccessToken(ctx context.Context, oauth2Conf oauth2.Config, authzCode string) (*oauth2.Token, error) {
	accessToken, err := oauth2Conf.Exchange(ctx, authzCode)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}
