package output

import (
	"fmt"
	"image"
	"net/http"

	"github.com/kingtingthegreat/ansi-converter/converter"
	"github.com/kingtingthegreat/top-fetch/config"
)

func UrlToImage(url string) (*image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get image from url")
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot decode image from response")
	}

	return &img, nil
}

func ImageToAnsi(img *image.Image) string {
	cfg := config.Config()

	return converter.Convert(*img, &cfg.ConverterConfig)
}
