package output

import (
	"log"
	"os"
	"sync"
)

func ReadBackup(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	return string(data)
}

func WriteBackup(file string, data string) *sync.WaitGroup {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		os.WriteFile(file, []byte(data), 0666)
		wg.Done()
	}()

	return &wg
}
