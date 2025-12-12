package fetch

import (
	"image"
)

func Fetch(web bool) (*image.Image, string, error) {
	var trackText string
	var img *image.Image
	var err error
	if web {
		trackText, img, err = WebFetch()
		if err != nil {
			return nil, "", err
		}
	} else {
		trackText, img, err = LocalFetch()
		if err != nil {
			return nil, "", err
		}
	}

	return img, trackText, nil
}
