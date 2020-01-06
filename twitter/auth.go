package twitter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//AuthConfig config
type AuthConfig struct {
	APIKey    string
	APISecret string
}

//AuthResponse response
type AuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

//GetBearerToken Get authorization token for twitter api request
func (config *AuthConfig) GetBearerToken() (string, error) {
	var apiKey = config.APIKey
	var apiSecret = config.APISecret
	var request *http.Request
	var response *http.Response
	var body []byte
	var authRes AuthResponse
	var authURL = "https://api.twitter.com/oauth2/token?grant_type=client_credentials"
	request, err := http.NewRequest("POST", authURL, nil)
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
