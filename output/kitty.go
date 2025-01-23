package output

import (
	"image"
	"os"
	"strings"

	"github.com/dolmen-go/kittyimg"
	"github.com/kingtingthegreat/top-fetch/config"
	"github.com/nfnt/resize"
)

func KittyOutput(trackText string, img *image.Image) {
	cfg := config.Config()

	os.Stdout.WriteString(strings.Repeat("\n", cfg.MarginTop+cfg.ConverterConfig.PaddingTop))

	os.Stdout.WriteString(strings.Repeat(" ", cfg.ConverterConfig.PaddingLeft))
	kittyimg.Fprint(os.Stdout, resize.Resize(uint(cfg.ConverterConfig.Dim), uint(cfg.ConverterConfig.Dim), *img, resize.Lanczos3))
	os.Stdout.WriteString(strings.Repeat(" ", cfg.ConverterConfig.PaddingRight))

	os.Stdout.WriteString(strings.Repeat("\n", cfg.ConverterConfig.PaddingBottom))
	os.Stdout.WriteString("\n" + trackText + "\n")
	os.Stdout.WriteString(strings.Repeat("\n", cfg.MarginBottom))
}
