package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"
)

/*
	Linker som kan være til hjelp.
	https://golang.org/pkg/bytes/#Count sjekk ut Repeat
	https://golang.org/pkg/strings/#Compare
	https://golang.org/pkg/unicode/utf8/#RuneCount

*/
func main() {

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	utf8.RuneCount(data)
	//

	bytes.Repeat(data, 5)

	var stringData = string(data)

	utf8.RuneCountInString(stringData)
	fmt.Println("bytes =", len(stringData))
	fmt.Println("runes =", utf8.RuneCountInString(stringData))
}
