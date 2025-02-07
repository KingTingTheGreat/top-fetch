package main

import (
	"log"
	"time"

	"github.com/kingtingthegreat/top-fetch/config"
	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/fetch"
	"github.com/kingtingthegreat/top-fetch/output"
)

func fetchAndDisplay(web bool) {
	var imageUrl, trackText string
	if web {
		imageUrl, trackText = fetch.WebFetch()
	} else {
		imageUrl, trackText = fetch.LocalFetch()
	}

	img, err := output.UrlToImage(imageUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	output.Output(img, trackText)
}

func main() {
	env.LoadEnv()
	config.ParseArgs()
	cfg := config.Config()

	if cfg.Timeout < 0 {
		// negative means no timeout
		fetchAndDisplay(cfg.Web)
	} else {
		timeout := time.Tick(time.Duration(cfg.Timeout) * time.Millisecond)
		done := make(chan bool)
		go func() { fetchAndDisplay(cfg.Web); done <- true }()
		select {
		case <-done:
			return
		case <-timeout:
			if cfg.Backup == "" {
				log.Fatal("Exceeded the ", cfg.Timeout, " millisecond time limit")
			}
			backupString := output.ReadBackup(cfg.Backup)
			if backupString == "" {
				log.Fatal("Exceeded the ", cfg.Timeout, " millisecond time limit and backup is empty")
			}
			output.OutputBackup(backupString)
		}
	}
}
