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
	fmt.Println("Navn:", filnavn)
	fmt.Println("Size: ", file.Size(), "bytes, ", "kibibytes, ", "mibibytes, ", "gibibytes")

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

//
// 	switch mode := file.Mode(); {
// 	case mode.IsDir():
// 			fmt.Println("directory", file.Mode().IsDir())
// 	case mode.IsRegular():
// 		fmt.Println("regular file", file.Mode().IsRegular())
// 	case mode&os.ModePerm != 0:
// 		fmt.Println("Has Unix permissions bits: ", file.Mode().Perm())
// 	case mode&os.ModeAppend != 0:
// 		fmt.Println("append only")
// 	case mode&os.ModeDevice != 0:
// 		fmt.Println("device file")
// 	case mode&os.ModeCharDevice != 0:
// 		fmt.Println("Unix character device")
// //	case ----
// //		fmt.Println("Unix block device")
// 	case mode&os.ModeSymlink != 0:
// 		fmt.Println("symbolic link")
// 	}
}
