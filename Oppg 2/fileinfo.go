package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"path/filepath"
)

func main() {
	filePtr := flag.String("f", "", "filnavn")
	flag.Parse()
	
	file, err := os.Stat(*filePtr)
	if err != nil {
		log.Fatal(err)
	}
	path, err := filepath.Abs(*filePtr)

	i64 := file.Size()
	b := float64(i64)
	kib := b / 1024
	mib := kib / 1024
	gib := mib / 1024

	fmt.Println("------------------------------------")
	fmt.Printf("Information about a file: %s \n", path)
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
}
