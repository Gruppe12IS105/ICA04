package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)
func main() {
	velgFil := os.Args[1]		//skriv inn: 'text1.txt' eller 'text2.txt'
	txtfil, err := ioutil.ReadFile(velgFil) //text1 og text2 blir []byte
	if err != nil {
		log.Fatal("\nError: Sjekk om argumentet er riktig!\n•Tips: argument skal være et filnavn (med endelse) \nog filen må ligge i samme pakke som go-filen\nEx: go run Oppg1.go text1.txt\n")
}
	if Carriage(txtfil) == true && LineFeed(txtfil) == true {
		fmt.Println("Både 0x0D (carriage return) og 0x0A (Line Feed) er brukt i filen: ", velgFil)
		} else if Carriage(txtfil) == true {
			fmt.Println("0x0D (Carriage return) er brukt i filen: ", velgFil)
		} else if LineFeed(txtfil) == true {
			fmt.Println("0x0A (Line Feed) er brukt i filen:", velgFil)
		}
	if Carriage(txtfil) == false && LineFeed(txtfil) == false {
			fmt.Println("Hverken 0x0D og 0x0A er brukt i teksten")
		} else if Carriage(txtfil) == false {
		fmt.Println("0x0D (carriage return) er ikke brukt i filen: ", velgFil)
		} else if LineFeed(txtfil) == false {
		fmt.Println("0x0A (Line Feed) er ikke brukt i filen:", velgFil)
	}
}
func Carriage(txtfil []byte) bool {
	for _, i := range txtfil {
		if i == 0x0D {
			return true
		}
	}
	return false
}
func LineFeed(txtfil []byte) bool {
	for _, i := range txtfil {
		if i == 0x0A {
			return true
		}
	}
	return false
}
