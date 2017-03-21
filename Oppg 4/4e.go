package main

import (
    "fmt"
    "os"
    "strings"
  )

func main() {
  fakultet := os.Args[1]
  m := make(map[string]string)
    m["Økonomi/samf.viten"] = "11"
    m["Lærer"] = "001"
    m["Kunstfag"] = "000"
    m["Humaniora/ped"] = "010"
    m["Helse/idrett"] = "011"
    m["Teknologi/realfag"] = "10"

if strings.HasPrefix(fakultet, "0") || strings.HasPrefix(fakultet, "1") {
  // DECODE
  reverse_m := make(map[string]string) // empty map
    for key, value := range m {
      reverse_m[value] = key
    }
  decode := reverse_m[fakultet]
    fmt.Println(fakultet, "har kode: ", decode)
} else {
  // ENCODE
  encode := m[fakultet]
    fmt.Println(fakultet, "har kode: ", encode)
  }
}
