package twitter

import (
	"os"
	"testing"
)

func TestGetBearerTokenSuccess(t *testing.T) {
	apiKey := os.Getenv("TWITTER_API_KEY")
	apiSecret := os.Getenv("TWITTER_API_SECRET")
	config := AuthConfig{APIKey: apiKey, APISecret: apiSecret}
	token, err := config.GetBearerToken()
	if err != nil {
		t.Fatal(err)
	}
	if len(token) <= 0 {
		t.Fatal("can not get token.")
	}
}
