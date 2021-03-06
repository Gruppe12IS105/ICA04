// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package diverse

import (
	"io"
	"log"
	"os"
)

// Hvordan kopiere filer
func main() {

	originalFile, err := os.Open("g.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	newFile, err := os.Create("h.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
