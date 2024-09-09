package maps

import (
	"log"
	"os"
)

func CountWords() {
	// Open a file
	f, err := os.Open("great-gatsby.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

}
