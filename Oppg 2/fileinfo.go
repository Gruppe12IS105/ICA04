package main

import (
	"fmt"
	"log"
	"os"
	"time"
)
//bare laste opp noe, bytter til en annen oppgave
//koden fungerer ikke

func main() {
	fi, err := os.Stat("fileinfo.go")
	if err != nil {
		log.Fatal(err)
	}

fmt.Printf("read %d bytes: %q\n", count, data[:count])
fmt.Printf("%s")
type FileInfo interface {
	Name() string
	Size([]byte) int64
	Mode() FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() interface{}
}
