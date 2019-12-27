package twitter

import (
	"net/http"
)

type AuthConfig struct {
	APIKey    string
	APISecret string
}

func (config *AuthConfig) getBearerToken() (string, error) {
	var apiKey = config.APIKey
	var apiSecret = config.APISecret
	var request *http.Request
	var response *http.Response
	var authUrl = "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
	request, err := http.NewRequest("POST", authUrl, nil)
	if err != nil {
		return "", err
	}
	request.SetBasicAuth(apiKey, apiSecret)
	client := &http.Client{}
	response, err = client.Do(request)
	if err != nil {
		return "", err
	}
}
