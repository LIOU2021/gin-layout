package helpers

import (
	"log"
	"os"
)

func OpenFile(fileName string) *os.File {
	fileExists := FileExists(fileName)

	if fileExists {
		f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		return f
	} else {
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		return f
	}
}
