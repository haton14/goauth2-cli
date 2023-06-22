package main

import (
	"golang.org/x/oauth2"
)

func main() {
	//ctx := context.Background()
	_ = &oauth2.Config{
		ClientID:     "client-id",
		ClientSecret: "client-secret",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/tokenn",
		},
		RedirectURL: "http://localhost:8080",
	}

}
