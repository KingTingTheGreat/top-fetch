package output

import (
	"fmt"
	"image"
	"net/http"

	"github.com/kingtingthegreat/ansi-converter/converter"
	"github.com/kingtingthegreat/top-fetch/config"
)

func UrlToAnsi(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to get image from url")
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot decode image from response")
	}

	cfg := config.Config()

	return converter.Convert(img, &cfg.ConverterConfig), nil
}
