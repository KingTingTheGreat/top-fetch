package spotify

import (
	"fmt"
	"io"
	"net/http"
)

// sends a request to the specified endpoint and returns the response bytes
func spotifyRequest(clientId, clientSecret, accessToken, refreshToken, endpoint string) ([]byte, string, error) {
	requestFunc := func(accessToken, endpoint string) ([]byte, error) {
		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			return []byte{}, err
		}

		req.Header.Set("Authorization", "Bearer "+accessToken)

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			return []byte{}, err
		}

		defer res.Body.Close()

		if res.StatusCode == http.StatusUnauthorized {
			return []byte{}, fmt.Errorf("unauthorized request")
		} else if res.StatusCode != http.StatusOK {
			return []byte{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return []byte{}, err
		}

		return body, nil
	}

	firstAttempt, err := requestFunc(accessToken, endpoint)
	if err == nil {
		return firstAttempt, "", nil
	}
	if err.Error() != "unauthorized request" {
		return []byte{}, "", err
	}

	newAccessToken, err := RefreshAccessToken(clientId, clientSecret, refreshToken)
	if err != nil {
		return []byte{}, "", fmt.Errorf("could not refresh access")
	}

	secondAttempt, err := requestFunc(newAccessToken, endpoint)
	return secondAttempt, newAccessToken, nil
}
