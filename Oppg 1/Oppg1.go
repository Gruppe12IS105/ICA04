package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func løkke() string {
	velgFil := os.Args[1]                   //skriv inn: 'text1.txt' eller 'text2.txt'
	txtfil, err := ioutil.ReadFile(velgFil) //text1 og 2 blir []byte
	if err != nil {
		return "Error"
	} else {
		for _, i := range txtfil {
			if i == 0x0D {
				return "ja"
			}
		}
	}
	return "nei"
}
func main() {
	løkke()
	if løkke() == "ja" {
		fmt.Println("0x0D (carriage return) er brukt i teksten")
	} else if løkke() == "nei" {
		fmt.Println("0x0D (carriage return) er ikke brukt i teksten")
	} else {
		fmt.Println("Error: Sjekk om argumentet er riktig!\n•Tips: argument skal være et filnavn (med endelse) \nog filen må ligge i samme pakke som go-filen\nEx: go run Oppg1.go text1.txt")
	}
}

//////////////////////////////////////////////////////////////////////////////

/*
Et førsteutkast - kan sjekke senere om jeg kan få den til å funke.
func main() {
	LesOgLøkke()
	fmt.Println()
}

func LesOgLøkke() bool {
	text1, err := ioutil.ReadFile("text1.txt") //text1 og 2 blir []byte
	if err != nil {
		fmt.Print(err)
	}
	for i := 0; i < len(text1); i++ {
		if text1[i] == 0x0D {
			fmt.Printf("%x er i teksten", text1[i])
		}
	}
	return
}

/* Oppg. 1 a)
	fmt.Println(text1)
	fmt.Printf("\n%X", text1)
	fmt.Println("\n----------------------------")

	b, err := ioutil.ReadFile("text2.txt")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(b)
	fmt.Printf("\n%X", b)
}
*/

/*
text1.txt bruker "carriage return" (hex: 0D/dec: 13)

fmt.Println(b) // Print ut som bytes (desimal)

str := string(b) // Konverter til string
fmt.Println(str) // Print som string
*/
