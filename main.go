package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"
)

// Command line arguments
var (
	filename string
)

type ngram string                   // A single n-gram
type probabilities map[rune]int     // The number of ocurances of each charachter
type ngrams map[ngram]probabilities // Map of n-grams to probobilites the the charachter folowing it

// getArgs stores command line arguments in global variables
func getArgs() {
	// Add flag deffenitions
	flag.StringVar(&filename, "f", "", "The filename to use as training data, one keysmash per line.")
	flag.Parse()

	if filename == "" {
		log.Fatal("no filename provided")
	}
}

// getFile returns a buffered reader from filename
func getFile(filename string) *bufio.Reader {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("could not open file:", filename, "error:", err)
	}

	// Wrap it in a buffer(4k size)
	fileBuf := bufio.NewReader(file)

	return fileBuf
}

func main() {
	getArgs() // Get cli args

	file := getFile(filename) // Get buffered reader

	// Test code
	text, err := file.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatalln("could not read file:", filename, "error:", err)
	}
	log.Println("file text:", text)

	//nGrams := parseNGrams(text)

	//newText := randSmash(ngrams)

	//fmt.Println(newText)
}
