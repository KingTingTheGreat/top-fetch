package main

import (
	"log"
	"time"

	"github.com/kingtingthegreat/top-fetch/config"
	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/fetch"
	"github.com/kingtingthegreat/top-fetch/output"
)

func fetchAndDisplay(web bool, backupFile string) {
	img, trackText, err := fetch.Fetch(web)
	if err != nil {
		backupString := output.ReadBackup(backupFile)
		if backupString == "" {
			log.Fatal("Something went wrong while fetching and backup is empty")
		}
		output.OutputBackup(backupString)
	}

	output.Output(img, trackText)
}

func fetchAndDisplayTimeout(web bool, done chan bool, backupFile string) {
	img, trackText, err := fetch.Fetch(web)
	if err != nil {
		done <- true
		backupString := output.ReadBackup(backupFile)
		if backupString == "" {
			log.Fatal("Something went wrong while fetching and backup is empty")
		}
		output.OutputBackup(backupString)
		done <- true
	}

	select {
	case <-done:
		return
	default:
		done <- true
		output.Output(img, trackText)
		done <- true
	}
}

func main() {
	env.LoadEnv()
	config.ParseArgs()
	cfg := config.Config()

	if cfg.Timeout < 0 {
		// negative means no timeout
		fetchAndDisplay(cfg.Web, cfg.Backup)
	} else {
		timeout := time.Tick(time.Duration(cfg.Timeout) * time.Millisecond)
		done := make(chan bool)
		go fetchAndDisplayTimeout(cfg.Web, done, cfg.Backup)
		select {
		case <-done:
			<-done
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
