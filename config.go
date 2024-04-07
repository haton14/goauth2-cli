package goauth2

import (
	"errors"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewOAuth2Config() (oauth2.Config, error) {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	redirectURI := os.Getenv("OAUTH_REDIRECT_URI")
	if clientID == "" {
		return oauth2.Config{}, errors.New("OAUTH_CLIENT_ID is not set")
	}
	if clientSecret == "" {
		return oauth2.Config{}, errors.New("OAUTH_CLIENT_SECRET is not set")
	}
	if redirectURI == "" {
		return oauth2.Config{}, errors.New("OAUTH_REDIRECT_URI is not set")
	}
	return oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"email", "profile", "https://www.googleapis.com/auth/photoslibrary.readonly", "openid"},
		RedirectURL:  redirectURI,
		Endpoint:     google.Endpoint,
	}, nil

}
