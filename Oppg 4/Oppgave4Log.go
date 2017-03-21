package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//1829, 1525, 420, 2166, 1506, 3093, 10 539
	fakultet := os.Args[1]
	fakInt, err := strconv.ParseFloat(fakultet, 64)
	if err != nil {
		fmt.Println(err)
	}
	total := 10539
	floatTotal := float64(total)
	fmt.Println(math.Log2(1 / (fakInt / floatTotal)))

	sum := math.Log2(1/(1829/floatTotal)) + math.Log2(1/(420/floatTotal)) + math.Log2(1/(2166/floatTotal)) + math.Log2(1/(1506/floatTotal)) + math.Log2(1/(3093/floatTotal)) + math.Log2(1/(1525/floatTotal))
	fmt.Println(sum)
}
