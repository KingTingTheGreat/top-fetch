package config

import (
	converter_config "github.com/kingtingthegreat/ansi-converter/config"
	"github.com/kingtingthegreat/ansi-converter/defaults"
)

var cfg config = config{
	Web:          false,
	Kitty:        false,
	Pix:          10,
	Path:         "source",
	File:         "",
	Wrap:         false,
	Timeout:      -1,
	MarginTop:    0,
	MarginRight:  0,
	MarginBottom: 0,
	MarginLeft:   0,
	ConverterConfig: converter_config.Config{
		Dim:           defaults.DEFAULT_DIM,
		Char:          defaults.DEFAULT_CHAR,
		FontRatio:     defaults.DEFAULT_RATIO,
		PaddingTop:    0,
		PaddingRight:  0,
		PaddingBottom: 0,
		PaddingLeft:   0,
	},
	Backup: "",
	Choice: 1,
}

func Config() *config {
	return &cfg
}
