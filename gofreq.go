/*
Gofreq counts frequencies for words in stdin or in the file.
Input: Text in an text file or in stdin.
Output: List of words and frequency of words in stdout.
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"go-freq/freq"
)

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

	// get sorted words by freq
	fw, err := freq.GetSortedFreq(string(text))
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	n := len(fw)
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%s,%d\n", fw[i].Key, fw[i].Value)
	}
}
