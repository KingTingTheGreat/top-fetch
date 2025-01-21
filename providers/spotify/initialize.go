package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func AuthUrl(clientId, redirectUri string) string {
	params := url.Values{}
	params.Set("client_id", clientId)
	params.Set("response_type", "code")
	params.Set("redirect_uri", redirectUri)
	params.Set("scope", "user-top-read user-read-email")

	return fmt.Sprintf("%s?%s", "https://accounts.spotify.com/authorize", params.Encode())
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func ExchangeCode(clientId, clientSecret, redirectUri, code string) (string, string, error) {
	params := url.Values{}
	params.Set("grant_type", "authorization_code")
	params.Set("code", code)
	params.Set("redirect_uri", redirectUri)

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(params.Encode()))
	if err != nil {
		return "", "", err
	}

	req.SetBasicAuth(clientId, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	defer res.Body.Close()

	// json.NewDecoder(res.Body).Decode(&var)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", "", err
	}

	return tokenResponse.AccessToken, tokenResponse.RefreshToken, nil

}
