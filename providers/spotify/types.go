package spotify

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
}

type Artist struct {
	Name string `json:"name"`
	Href string `json:"href"`
	Uri  string `json:"uri"`
}

type Album struct {
	Name        string   `json:"name"`
	Href        string   `json:"href"`
	Uri         string   `json:"uri"`
	Artists     []Artist `json:"artists"`
	ReleaseDate string   `json:"release_date"`
	TotalTracks int      `json:"total_tracks"`
	Images      []Image  `json:"images"`
}

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type Track struct {
	Name         string       `json:"name"`
	Album        Album        `json:"album"`
	Artists      []Artist     `json:"artists"`
	Uri          string       `json:"uri"`
	Href         string       `json:"href"`
	TrackNumber  int          `json:"track_number"`
	Popularity   int          `json:"popularity"`
	ExternalUrls ExternalUrls `json:"external_urls"`
}

type ProfileResponse struct {
	DisplayName string `json:"display_name"`
	Href        string `json:"href"`
	SpotifyId   string `json:"id"`
}

type TopTracksResponse struct {
	Items []Track `json:"items"`
	Total int     `json:"total"`
	Limit int     `json:"limit"`
}
