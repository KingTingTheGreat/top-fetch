package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kingtingthegreat/top-fetch/config"
	"github.com/kingtingthegreat/top-fetch/env"
	"github.com/kingtingthegreat/top-fetch/output"
)

func fetchAndDisplay(web bool) {
	var imageUrl, trackText string

	// if web {
	// 	imageUrl, trackText = fetch.WebFetch()
	// } else {
	// 	imageUrl, trackText = fetch.LocalFetch()
	// }
	// log.Println("converting")
	ansiImage, err := output.UrlToAnsi(imageUrl)
	if err != nil {
		log.Fatal(err)
	}

	output.Output(ansiImage, trackText)
}

func main() {
	start := time.Now()

	env.LoadEnv()
	config.ParseArgs()
	cfg := config.Config()

	fmt.Println("hi from top fetch lalala", start)
	if cfg.Timeout < 0 {
		// negative means no timeout
		fmt.Println(cfg.Web)
	} else {
		c := make(chan bool)
		go func() { fmt.Println(cfg.Web); c <- true }()
		for time.Now().Before(start.Add(time.Duration(cfg.Timeout) * time.Millisecond)) {
			select {
			case <-c:
				return
			default:
				continue
			}
		}
		log.Fatal("Exceeded the ", cfg.Timeout, " millisecond time limit")
	}

}
