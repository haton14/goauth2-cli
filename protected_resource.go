package goauth2

import (
	"context"
	"io"

	"golang.org/x/oauth2"
)

func GetProtectedResource(ctx context.Context, oauth2Conf oauth2.Config, accessToken *oauth2.Token) (string, error) {
	oauth2Client := oauth2Conf.Client(ctx, accessToken)
	resp, err := oauth2Client.Get("https://photoslibrary.googleapis.com/v1/mediaItems")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
