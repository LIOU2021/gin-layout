package helpers

import "os"

func OpenFile(fileName string) *os.File {
	_, error := os.Stat(fileName)

	if os.IsNotExist(error) {
		f, _ := os.Create(fileName)
		return f
	} else {
		f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		return f
	}
}
