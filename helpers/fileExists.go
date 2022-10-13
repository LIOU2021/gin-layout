package helpers

import "os"

func FileExists(fileName string) bool {
	_, error := os.Stat(fileName)

	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}
