package fetch

import (
	"image"
)

func Fetch(web bool) (*image.Image, string, error) {
	var imageUrl, trackText string
	var err error
	if web {
		imageUrl, trackText, err = WebFetch()
		if err != nil {
			return nil, "", err
		}
	} else {
		imageUrl, trackText, err = LocalFetch()
		if err != nil {
			return nil, "", err
		}
	}

	img, err := UrlToImage(imageUrl)
	if err != nil {
		return nil, "", err
	}

	return img, trackText, nil
}
