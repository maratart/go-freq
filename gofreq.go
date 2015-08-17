/*
Gofreq counts frequencies for words in stdin.
Input: Text in an text file or in stdin
Output: List of words and frequency of words in stdout.
*/

package main
import (
	"fmt"
	"strings"
	"unicode"
	"sort"
	"io"
	"os"
	"io/ioutil"
	"log"
)

// Map of words to counts
type wordsFreq map[string]int

// freq calculates count of the words in the text
func freq(text string) (wordsFreq, error) {

	wf := make(wordsFreq)

	// split text into words
	f := func(c rune) bool { return !unicode.IsLetter(c) }
	words := strings.FieldsFunc(text, f)

	// words to lower
	for i, _ := range words {
		words[i] = strings.ToLower(words[i])
	}

	// calculate freq
	for i, _ := range words {
		wf[words[i]] = wf[words[i]] + 1
	}

	return wf, nil
}

// A data structure to hold a key/value pair.
type Pair struct {
	Key string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

func openStdinOrFile() io.Reader {
	var err error
	r := os.Stdin
	if len(os.Args) > 1 {
		r, err = os.Open(os.Args[1])
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	return r
}

func main() {

	r := openStdinOrFile()
	text, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// get freq of all words
	wf, _ := freq(string(text))

	// sort words by freq
	fw := sortMapByValue(wf)
	n := len(fw)
	for i := n-1; i >= 0; i-- {
		fmt.Printf("%s,%d\n", fw[i].Key, fw[i].Value)
	}
}
