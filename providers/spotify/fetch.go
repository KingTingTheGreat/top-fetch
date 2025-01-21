package spotify

import (
	"encoding/json"
	"fmt"
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
func GetUserTopTrack(clientId, clientSecret, accessToken, refreshToken string) (*Track, string, error) {
	body, newAccessToken, err := spotifyRequest(clientId, clientSecret, accessToken, refreshToken, "https://api.spotify.com/v1/me/top/tracks?time_range=short_term&limit=1")
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

	return &topTracksResponse.Items[0], newAccessToken, nil

}
