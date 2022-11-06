package helpers

import (
	"fmt"
	"log"
)

func Logger(text string) {
	f := OpenFile("./logs/debug.log")

	defer f.Close()

	_, err2 := f.WriteString(text + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
