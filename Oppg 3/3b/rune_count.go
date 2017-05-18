// https://www.daniweb.com/programming/computer-science/code/496353/sort-a-word-frequency-count-golang
// Kode fra linken ovenfor er delvis gjenbrukt her

package main

import (
	"fmt"
	"regexp"
	"sort"
	"io/ioutil"
	"os"
)

func word_count(words []string) map[string]int {
	word_freq := make(map[string]int)
	for _, word := range words {
		word_freq[word]++
	}
	return word_freq
}

type word_struct struct {
	freq int
	word string
}

// word_struct will be displayed in this format
func (p word_struct) String() string {
	return fmt.Sprintf("%3d   %s", p.freq, p.word)
}

// by_freq implements sort.Interface for []word_struct based on the freq field
type by_freq []word_struct
func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].freq < a[j].freq }

// by_word implements sort.Interface for []word_struct based on the word field
type by_word []word_struct
func (a by_word) Len() int           { return len(a) }
func (a by_word) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_word) Less(i, j int) bool { return a[i].word < a[j].word }


func main() {
	file, err := ioutil.ReadFile(os.Args[1]) // Open input file from terminal to analyze
	if err != nil {
		fmt.Print(err)
	}
	text := string(file) 									// convert content to a 'string'
	re := regexp.MustCompile("\\w") 			// matches a character, removes any punctuation/whitespace
	words := re.FindAllString(text, -1)
	word_map := word_count(words)					// word_map() returns a map of word:frequency pairs

	// convert the map to a slice of structures for sorting
	// create pointer to an instance of word_struct
	pws := new(word_struct)
	struct_slice := make([]word_struct, len(word_map))
	ix := 0
	for key, val := range word_map {
		pws.freq = val
		pws.word = key
		struct_slice[ix] = *pws
		ix++
	}

	fmt.Println("Characters sorted by frequency")
	fmt.Println("Highest--Lowest:")
	fmt.Println("------------------------------")
	sort.Sort(sort.Reverse(by_freq(struct_slice)))
	for ix := 0; ix < 5; ix++ {
		fmt.Printf("%v\n", struct_slice[ix])
	}
	fmt.Println("------------------------------")
}
