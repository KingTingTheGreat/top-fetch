package config

import (
	converter_config "github.com/kingtingthegreat/ansi-converter/config"
)

type config struct {
	Web                 bool
	Kitty               bool
	Pix                 float64
	TopFetchId          string
	SpotifyClientId     string
	SpotifyClientSecret string
	SpotifyAccessToken  string
	SpotifyRefreshToken string
	Path                string
	File                string
	Wrap                bool
	Timeout             int
	Silent              bool
	MarginTop           int
	MarginRight         int
	MarginBottom        int
	MarginLeft          int
	ConverterConfig     converter_config.Config
	Env                 string
	Backup              string
}

const (
	WEB            = "web"
	ID             = "id"
	KITTY          = "kitty"
	PIX            = "pix"
	DIM            = "dim"
	CHAR           = "char"
	RATIO          = "ratio"
	PATH           = "path"
	FILE           = "file"
	TIMEOUT        = "timeout"
	SILENT         = "silent"
	PADDING        = "p"
	PADDING_TOP    = "pT"
	PADDING_RIGHT  = "pR"
	PADDING_BOTTOM = "pB"
	PADDING_LEFT   = "pL"
	MARGIN         = "m"
	MARGIN_TOP     = "mT"
	MARGIN_RIGHT   = "mR"
	MARGIN_BOTTOM  = "mB"
	MARGIN_LEFT    = "mL"
	ENV            = "env"
	BACKUP         = "backup"
)

const WRAP = "wrap"
