// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("files/pg100.txt")
	if err != nil {
		log.Fatal(err)
	}
	buffRead := bufio.NewReader(file)

	byteSlice := make([]byte, 4046)
	byteSlice, err = buffRead.Peek(4046)
	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := buffRead.Buffered()
	//Antall linjer
	nl := 0
	for i := 1; i <= unflushedBufferSize; i++ {
		c := (bytes.Count(byteSlice, []byte("\x0A")))
		nl = nl + c
	}
	fmt.Println("Number of lines: ", nl)
}
