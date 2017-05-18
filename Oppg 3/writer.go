package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

//eksempler av flere typer write funksjoner.
//Funksjonene skriver i write.txt filen, og ikke de i "files" mappen
func main() {
	write()
	quickWrite()
	bufferedWrite()
}

func write() {
	//opens the write.txt file
	file, err := os.OpenFile(
		"write.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Writes the bytes
	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Write Wrote %d bytes. \n", bytesWritten)
}

func quickWrite() {
	err := ioutil.WriteFile("quickWrite.txt", []byte("This is quickwrite!\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func bufferedWrite() {
	file, err := os.OpenFile("bufferedWrite.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	bytesWritten, err := bufferedWriter.Write(
		[]byte{65, 66, 67},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("buffered Write: Bytes written: %d\n", bytesWritten)

	bytesWritten, err = bufferedWriter.WriteString(
		"Buffered String\n",
	)

	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("buffered Write: Bytes buffered: %d\n", unflushedBufferSize)

	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("buffered Write: Available buffer: %d\n", bytesAvailable)

	bufferedWriter.Flush()

	bufferedWriter.Reset(bufferedWriter)

}
