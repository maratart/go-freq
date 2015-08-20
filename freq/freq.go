/*
Gofreq counts frequencies for words in stdin or in the file.
Input: Text in an text file or in stdin.
Output: List of words and frequency of words in stdout.
*/

package freq

import (
	"log"
	"sort"
	"strings"
	"unicode"
)

// WordsFreq is a map of words to counts
type WordsFreq map[string]int

// Freq calculates count of the words in the text
func Freq(text string) (WordsFreq, error) {
	wf := make(WordsFreq)

	// split text into words
	f := func(c rune) bool { return !unicode.IsLetter(c) }
	words := strings.FieldsFunc(text, f)

	// words to lower
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	// calculate freq
	for i := range words {
		wf[words[i]] = wf[words[i]] + 1
	}

	return wf, nil
}

// Pair is a data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// PairList is a slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
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

// GetSortedFreq takes text and returns sorted list of words with frequencies
func GetSortedFreq(text string) (PairList, error) {

	// Get freq
	wf, err := Freq(string(text))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Sort freq
	fw := sortMapByValue(wf)
	return fw, nil
}
