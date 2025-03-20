package output

import (
	"image"

	"github.com/kingtingthegreat/ansi-converter/converter"
	"github.com/kingtingthegreat/top-fetch/config"
)

func ImageToAnsi(img *image.Image) string {
	cfg := config.Config()

	return converter.Convert(*img, &cfg.ConverterConfig)
}
