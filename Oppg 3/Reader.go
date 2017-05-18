package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	readUpTo()
	readAtLeast()
	readAll()
	readWholeToMem()
	readBuff()
	readScan()
}

func readUpTo() {
	file, err := os.Open("files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}

	bSlice := make([]byte, 30)
	bRead, err := file.Read(bSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Up to %d bytes read: \n %s", bRead, bSlice)
}

func readAtLeast() {
	file, err := os.Open("files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	alSlice := make([]byte, 60)
	minB := 50
	bReadA, err := io.ReadAtLeast(file, alSlice, minB)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("At least %d bytes read: \n %s", bReadA, alSlice)
}

func readAll() {
	file, err := os.Open("files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	text, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of bytes read:", len(text))
	fmt.Printf("Data as string: %s\n", text)
}

func readWholeToMem() {
	toxt, err := ioutil.ReadFile("files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read to memory: %s\n", toxt)
}

func readBuff() {
	tixt, err := os.Open("files/text1.txt")
	if err != nil {
		log.Fatal(err)
	}
	boffReader := bufio.NewReader(tixt)
	dotoBotes, err := boffReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read buffered bytes: %s\n", dotoBotes)
	dotoStronk, err := boffReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read buffered string: %s\n", dotoStronk)
}

func readScan() {
	// Åpner en fil. hvis ikke verdien er null = error
	file, err := os.Open("g.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Lager en scanner
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	success := scanner.Scan()
	if success == false {

		err = scanner.Err()
		if err == nil {
			log.Println("Scan completed and reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// Get data from scan with Bytes() or Text()
	fmt.Println("First word found by scanner", scanner.Text())
}
