package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filnavn := os.Args[1]
	file, err := os.Stat(filnavn)
	if err != nil {
		log.Fatal(err)
	}

	/*
		En litt midlertidig løsning til vi får
		import"./github.com/alecthomas/units" til å funke.
		Har kopiert pakken units slik at den blir importert,
		men metodene kan ikke brukes til tross for at de er public
	*/
	i64 := file.Size()
	b := float64(i64)
	kib := b / 1024
	mib := kib / 1024
	gib := mib / 1024

	/*
		file.Size() blir ikke alltid angitt i bytes.
		Ergo så dette kan være en potensiell bug, men har testet den
		på filer som tar flere Kb, og har vært riktig så langt.
		**Dersom ikke file.Size() er korrekt vil ikke resten (kib, mib og gib)
		være det heller!!!**
	*/

	/*
		På .app-filer står de listet opp som "Directory" (og generelt alle filer
		med "Show package content" (valg ved høyreklikk på macOS))
		som tar 105 bytes. Finner ingen filer i contents
		som tar 105 bytes heller.
	*/
	fmt.Println("------------------------------------")
	fmt.Println("Information about a file:", filnavn)
	fmt.Println("Size:", b, "bytes,", kib, "kibibytes,", mib, "mibibytes,", gib, "gibibytes")
	if file.Mode().IsDir() == true {
		fmt.Println("Is a directory")
	} else if file.Mode().IsDir() == false {
		fmt.Println("Is not a directory")
	}
	if file.Mode().IsRegular() == true {
		fmt.Println("Is a regular file")
	} else if file.Mode().IsRegular() == false {
		fmt.Println("Is not a regular file")
	}
	fmt.Println("Has Unix permissions bits: ", file.Mode().Perm())
	if file.Mode()&os.ModeAppend == os.ModeAppend {
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}
	if file.Mode()&os.ModeDevice == os.ModeDevice {
		fmt.Println("Is a device file")
		if file.Mode()&os.ModeCharDevice == os.ModeCharDevice {
			fmt.Println("Is a Unix character device")
			fmt.Println("Is not a Unix block device")
		} else {
			fmt.Println("Is not a Unix character device")
			fmt.Println("Is a Unix block device")
		}
	} else {
		fmt.Println("Is not a device file")
		fmt.Println("Is not a Unix character device")
		fmt.Println("Is not a Unix block device")
	}
	if file.Mode()&os.ModeSymlink == os.ModeSymlink {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}
	fmt.Println("------------------------------------")

	///TEST///////////////////////////
	fmt.Println("------Vil bli samme utskrift, men er litt alternativ koding----------")
	fmt.Println("File name:", filnavn)
	fmt.Println("Size:", b, "bytes,", kib, "kibibytes,", mib, "mibibytes,", gib, "gibibytes")
	if file.Mode().IsDir() == true {
		fmt.Println("Is a directory")
	} else if file.Mode().IsDir() == false {
		fmt.Println("Is not a directory")
	}
	if file.Mode().IsRegular() == true {
		fmt.Println("Is a regular file")
	} else if file.Mode().IsRegular() == false {
		fmt.Println("Is not a regular file")
	}
	fmt.Println("Has Unix permissions bits: ", file.Mode().Perm())
	if file.Mode()&os.ModeAppend != 0 {
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}
	if file.Mode()&os.ModeDevice != 0 {
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Is not a device file")
	}
	if file.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("Is a Unix character device")
	} else {
		fmt.Println("Is not a Unix character device")
	}
	//	if file.Mode()
	//		fmt.Println("Unix block device")
	if file.Mode()&os.ModeSymlink != 0 {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}
	fmt.Println("------------------------------------")
}
