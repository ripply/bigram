package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	argsCount := len(os.Args)

	if argsCount <= 1 {
		fmt.Println("Usage: ./server file-to-read")
		example := "the quick brown fox and the quick blue hare."

		bigrams := computeBigram(example)
		printBigram(bigrams)
	} else {
		// load file
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(
				fmt.Errorf(
					"Failed to read file '%s': %s",
					os.Args[1],
					err,
				),
			)
		}

		b, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(
				fmt.Errorf(
					"Failed to fully read file: %s",
					err,
				),
			)
		}

		bigrams := computeBigram(string(b))
		printBigram(bigrams)
	}
}

func printBigram(bigrams map[string]int) {
	for bigram, count := range bigrams {
		fmt.Printf("\"%s\" %d\n", bigram, count)
	}
}

func computeBigram(input string) map[string]int {
	input = strings.ToLower(input)

	excludedCharacters := []string{"."}

	bigrams := make(map[string]int)

	for _, excludedCharacter := range excludedCharacters {
		input = strings.Replace(input, excludedCharacter, "", -1)
	}

	words := strings.Split(input, " ")
	wordCount := len(words)

	for i := 0; i < wordCount; i++ {
		if i+1 == wordCount {
			continue
		}

		bigram := fmt.Sprintf("%s %s", words[i], words[i+1])
		bigrams[bigram]++
	}

	return bigrams
}
