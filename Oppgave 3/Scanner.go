package main

import (
    "os"
    "log"
    "fmt"
    "bufio"
)

func main() {
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
    fmt.Println("Første ord som ble funnet", scanner.Text())

    // Call scanner.Scan() again to find next token
}
