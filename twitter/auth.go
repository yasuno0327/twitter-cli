package twitter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type AuthConfig struct {
	APIKey    string
	APISecret string
}

type AuthResponse struct {
	TokenType   string `"json":"token_type"`
	AccessToken string `"json":"access_token"`
}

func (config *AuthConfig) getBearerToken() (string, error) {
	var apiKey = config.APIKey
	var apiSecret = config.APISecret
	var request *http.Request
	var response *http.Response
	var body []byte
	var authRes AuthResponse
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
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &authRes)
	if err != nil {
		return "", err
	}
	return authRes.AccessToken, nil
}
