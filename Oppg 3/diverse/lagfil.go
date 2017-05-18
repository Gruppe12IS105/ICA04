// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package diverse

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

// Lager nytt txt av g. Hvis verdien er null , kraskjer den.
func main() {
	newFile, err = os.Create("g.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}
