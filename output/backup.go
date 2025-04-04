package output

import (
	"log"
	"os"
	"sync"
)

func ReadBackup(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func WriteBackup(file string, data string, wg *sync.WaitGroup) {
	os.WriteFile(file, []byte(data), 0666)
	wg.Done()
}
