package output

import (
	"image"
	"os"
	"strings"
	"sync"

	"github.com/dolmen-go/kittyimg"
	"github.com/kingtingthegreat/top-fetch/config"
	"github.com/nfnt/resize"
)

func KittyOutput(trackText string, img *image.Image) {
	cfg := config.Config()

	var builder strings.Builder

	builder.WriteString(strings.Repeat("\n", cfg.MarginTop+cfg.ConverterConfig.PaddingTop))

	builder.WriteString(strings.Repeat(" ", cfg.ConverterConfig.PaddingLeft))
	kittyimg.Fprint(&builder, resize.Resize(uint(cfg.ConverterConfig.Dim), uint(cfg.ConverterConfig.Dim), *img, resize.Lanczos3))
	builder.WriteString(strings.Repeat(" ", cfg.ConverterConfig.PaddingRight))

	builder.WriteString(strings.Repeat("\n", cfg.ConverterConfig.PaddingBottom))
	builder.WriteString("\n" + trackText + "\n")
	builder.WriteString(strings.Repeat("\n", cfg.MarginBottom))

	var wg *sync.WaitGroup
	if cfg.Backup != "" {
		wg = WriteBackup(cfg.Backup, builder.String())
	}

	os.Stdout.WriteString(builder.String())
	wg.Wait()
}
