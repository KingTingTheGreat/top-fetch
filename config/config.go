package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	converter_config "github.com/kingtingthegreat/ansi-converter/config"
	"github.com/kingtingthegreat/ansi-converter/defaults"
	"github.com/kingtingthegreat/top-fetch/env"
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
)

const WRAP = "wrap"

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
}

func ParseArgs() {
	// configure args
	for _, argValStr := range os.Args[1:] {
		argVal := strings.SplitN(argValStr, "=", 2)
		arg := argVal[0]
		val := ""
		if len(argVal) > 1 {
			val = argVal[1]
		}
		switch arg {
		case WEB:
			cfg.Web = true
			if val != "" {
				cfg.TopFetchId = val
			}
		case ID:
			cfg.Web = true
			if val != "" {
				cfg.TopFetchId = val
			}
		case KITTY:
			cfg.Kitty = true
		case PIX:
			newPix, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal("invalid px per char")
			}
			cfg.Pix = newPix
		case DIM:
			newDim, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal("invalid dim")
			}
			cfg.ConverterConfig.Dim = newDim
		case CHAR:
			cfg.ConverterConfig.Char = val
		case RATIO:
			newRatio, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal("invalid ratio")
			}
			cfg.ConverterConfig.FontRatio = newRatio
		case FILE:
			cfg.File = val
		case TIMEOUT:
			newTimeout, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid timeout")
			}
			cfg.Timeout = newTimeout
		case SILENT:
			cfg.Silent = true
		case PADDING:
			newPadding, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid padding")
			}
			cfg.ConverterConfig.PaddingTop = newPadding
			cfg.ConverterConfig.PaddingRight = newPadding
			cfg.ConverterConfig.PaddingBottom = newPadding
			cfg.ConverterConfig.PaddingLeft = newPadding
		case PADDING_TOP:
			newPaddingTop, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid padding top")
			}
			cfg.ConverterConfig.PaddingTop = newPaddingTop
		case PADDING_RIGHT:
			newPaddingRight, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid padding right")
			}
			cfg.ConverterConfig.PaddingRight = newPaddingRight
		case PADDING_BOTTOM:
			newPaddingBottom, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid padding bottom")
			}
			cfg.ConverterConfig.PaddingBottom = newPaddingBottom
		case PADDING_LEFT:
			newPaddingLeft, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid padding left")
			}
			cfg.ConverterConfig.PaddingLeft = newPaddingLeft
		case MARGIN:
			newMargin, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid margin")
			}
			cfg.MarginTop = newMargin
			cfg.MarginRight = newMargin
			cfg.MarginBottom = newMargin
			cfg.MarginLeft = newMargin
		case MARGIN_TOP:
			newMarginTop, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid margin top")
			}
			cfg.MarginTop = newMarginTop
		case MARGIN_RIGHT:
			newMarginRight, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid margin right")
			}
			cfg.MarginRight = newMarginRight
		case MARGIN_BOTTOM:
			newMarginBottom, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid margin bottom")
			}
			cfg.MarginBottom = newMarginBottom
		case MARGIN_LEFT:
			newMarginLeft, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal("invalid margin left")
			}
			cfg.MarginLeft = newMarginLeft
		case ENV:
			if strings.HasPrefix(val, "~") {
				home, err := os.UserHomeDir()
				if err != nil {
					log.Fatal("could not evaluate home directory")
				}
				val, _ = strings.CutPrefix(val, "~")
				cfg.Env = filepath.Join(home, val)
			} else {
				cfg.Env = val
			}
		}
	}
	cfg.ConverterConfig.PaddingRight += cfg.MarginRight
	cfg.ConverterConfig.PaddingLeft += cfg.MarginLeft

	// prioritize env file
	if cfg.Env != "" {
		envMap, err := godotenv.Read(cfg.Env)
		if err != nil {
			log.Fatal(err)
		}

		if cfg.TopFetchId == "" {
			cfg.TopFetchId = envMap["TOP_FETCH_ID"]
		}
		cfg.SpotifyClientId = envMap["SPOTIFY_CLIENT_ID"]
		cfg.SpotifyClientSecret = envMap["SPOTIFY_CLIENT_SECRET"]
		cfg.SpotifyAccessToken = envMap["SPOTIFY_ACCESS_TOKEN"]
		cfg.SpotifyRefreshToken = envMap["SPOTIFY_REFRESH_TOKEN"]
	}

	// next try compiled env file, then environment variables
	if cfg.TopFetchId == "" {
		cfg.TopFetchId = env.EnvVal("TOP_FETCH_ID")
	}
	if cfg.SpotifyClientId == "" {
		cfg.SpotifyClientId = env.EnvVal("SPOTIFY_CLIENT_ID")
	}
	if cfg.SpotifyClientSecret == "" {
		cfg.SpotifyClientSecret = env.EnvVal("SPOTIFY_CLIENT_SECRET")
	}
	if cfg.SpotifyAccessToken == "" {
		cfg.SpotifyAccessToken = env.EnvVal("SPOTIFY_ACCESS_TOKEN")
	}
	if cfg.SpotifyRefreshToken == "" {
		cfg.SpotifyRefreshToken = env.EnvVal("SPOTIFY_REFRESH_TOKEN")
	}
}

func Config() *config {
	return &cfg
}
