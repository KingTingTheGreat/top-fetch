package fetch

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"

	"github.com/kingtingthegreat/top-fetch/config"
)

func WebFetch() (string, *image.Image, error) {
	cfg := config.Config()
	if cfg.TopFetchId == "" {
		log.Fatal("TopFetch id is not set")
	}

	res, err := http.Get(fmt.Sprintf("https://top-fetch.jting.org/track?id=%s&choice=%d", cfg.TopFetchId, cfg.Choice))
	if err != nil {
		return "", nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	dataList := bytes.SplitN(bodyBytes, []byte("\x1d"), 2)
	if len(dataList) < 2 {
		return "", nil, fmt.Errorf("invalid response format from Top Fetch server")
	}

	text := string(dataList[0])

	imgBytes := dataList[1]
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	// img, _, err := image.Decode(res.Body)
	if err != nil {
		panic(err)
	}

	return text, &img, nil
}
