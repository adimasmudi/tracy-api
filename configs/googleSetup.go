package configs

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleOAuthConfig() *oauth2.Config {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  os.Getenv("redirect_url"),
		ClientID:     os.Getenv("client_id"),
		ClientSecret: os.Getenv("client_secret"),
		Scopes:       []string{"email"},
		Endpoint:     google.Endpoint,
	}

	return googleOauthConfig
}