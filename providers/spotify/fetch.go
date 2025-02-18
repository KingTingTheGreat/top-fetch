package spotify

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
)

// Takes in AccessToken and RefreshToken.
// Returns user's SpotifyId.
// If the provided AccessToken has expired, return the new AccessToken.
func GetSpotifyId(clientId, clientSecret, accessToken, refreshToken string) (string, string, error) {
	body, newAccessToken, err := spotifyRequest(clientId, clientSecret, accessToken, refreshToken, "https://api.spotify.com/v1/me")
	if err != nil {
		return "", "", err
	}

	var userProfileRes ProfileResponse
	if err := json.Unmarshal(body, &userProfileRes); err != nil {
		return "", "", err
	}

	return userProfileRes.SpotifyId, newAccessToken, nil
}

// Takes in AccessToken and RefreshToken.
// Returns user's top song of type Track.
// If the provided AccessToken has expired, return the new AccessToken.
func GetUserTopTrack(clientId, clientSecret, accessToken, refreshToken string, choice int) (*Track, string, error) {
	body, newAccessToken, err := spotifyRequest(clientId, clientSecret, accessToken, refreshToken, fmt.Sprintf("https://api.spotify.com/v1/me/top/tracks?time_range=short_term&limit=%d", choice))
	if err != nil {
		return nil, "", err
	}

	var topTracksResponse TopTracksResponse
	if err := json.Unmarshal(body, &topTracksResponse); err != nil {
		return nil, "", err
	}

	if len(topTracksResponse.Items) == 0 {
		return nil, "", fmt.Errorf("no top track found")
	}

	max := big.NewInt(int64(len(topTracksResponse.Items)))
	randIndex, err := rand.Int(rand.Reader, max)
	if err != nil {
		randIndex = big.NewInt(0)
	}

	return &topTracksResponse.Items[randIndex.Int64()], newAccessToken, nil

}
