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
	i64 := file.Size()
	b := float64(i64)
	kib := b / 1024
	mib := kib / 1024
	gib := mib / 1024

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
	} else {
		fmt.Println("Is not a device file")
	}
	if file.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		fmt.Println("Is a Unix character device")
	} else {
		fmt.Println("Is not a Unix character device")
	}
//	if file.Mode()
//		fmt.Println("Unix block device")
	if file.Mode()&os.ModeSymlink == os.ModeSymlink {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}
	fmt.Println("------------------------------------")

// TEST
	fmt.Println("------TEST--------TEST-------TEST----------")
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
