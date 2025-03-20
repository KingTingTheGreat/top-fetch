package fetch

import (
	"image"
)

func Fetch(web bool) (*image.Image, string, error) {
	var imageUrl, trackText string
	if web {
		imageUrl, trackText = WebFetch()
	} else {
		imageUrl, trackText = LocalFetch()
	}

	img, err := UrlToImage(imageUrl)
	if err != nil {
		return nil, "", err
	}

	return img, trackText, nil
}
