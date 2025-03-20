package output

import (
	"image"
	"log"
	"os"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/kingtingthegreat/top-fetch/config"
)

func centerTrackText(trackText string, dim int, left, right int) string {
	length := utf8.RuneCountInString(strings.Split(trackText, "\x1B")[0] + "🎵")

	if length > dim {
		return trackText
	}

	width := left + dim + right
	rightPad := (width - length) / 2
	leftPad := width - length - rightPad
	return strings.Repeat(" ", leftPad) + trackText + strings.Repeat(" ", rightPad)
}

func Output(img *image.Image, trackText string) {
	cfg := config.Config()

	// display image itself
	if cfg.Kitty {
		trackText = centerTrackText(trackText, int(cfg.ConverterConfig.Dim/cfg.Pix),
			cfg.ConverterConfig.PaddingLeft, cfg.ConverterConfig.PaddingRight)
		KittyOutput(trackText, img)
		return
	}

	// display ansi image
	trackText = centerTrackText(trackText, int(cfg.ConverterConfig.Dim),
		cfg.ConverterConfig.PaddingLeft, cfg.ConverterConfig.PaddingRight)
	ansiImage := ImageToAnsi(img)
	outputString := strings.Repeat("\n", cfg.MarginTop) + ansiImage + "\n" + trackText + "\n" + strings.Repeat("\n", cfg.MarginBottom)

	wg := sync.WaitGroup{}
	if cfg.Backup != "" {
		wg.Add(1)
		go WriteBackup(cfg.Backup, outputString, &wg)
	}

	// write to desired output
	if cfg.File != "" {
		outputFile, err := WriteToFile(outputString)
		if err != nil {
			log.Fatal(err.Error())
		}
		os.Stdout.WriteString(outputFile)
	} else {
		os.Stdout.WriteString(outputString)
	}
	wg.Wait()
}

func OutputBackup(backupString string) {
	cfg := config.Config()

	if cfg.Kitty || cfg.File == "" {
		os.Stdout.WriteString(backupString)
	} else {
		outputFile, err := WriteToFile(backupString)
		if err != nil {
			log.Fatal(err.Error())
		}
		os.Stdout.WriteString(outputFile)
	}
}
