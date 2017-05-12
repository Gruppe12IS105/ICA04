package main

import (
    "os"
    "fmt"
    "log"
)

// Posisjonen etter seek burde være 0,0: 9

func main() {
    file, _ := os.Open("g.txt")
    defer file.Close()

    var offset int64 = 15

    var whence int = 5
    newPosition, err := file.Seek(offset, whence)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Just moved to 5:", newPosition)


    newPosition, err = file.Seek(-11, 2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Just moved back two:", newPosition)

    currentPosition, err := file.Seek(0, 1)
    fmt.Println("Current position:", currentPosition)

    newPosition, err = file.Seek(9, 0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Position after seeking 0,0:", newPosition)
}
