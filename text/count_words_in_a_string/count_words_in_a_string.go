/**
Counts the number of individual words in a string. For added complexity read these strings in from a text file and generate a summary.
*/
package count_words_in_a_string

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func CountWordsInAString() {
	data, err := ioutil.ReadFile("text/count_words_in_a_string/long_text.txt")
	if err != nil {
		panic(err)
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9\\s]+")
	if err != nil {
		panic(err)
	}

	content := strings.ToLower(reg.ReplaceAllString(string(data), ""))
	words := strings.Split(content, " ")

	wordsMap := make(map[string]int)

	var mostUserWord string
	countOfMostUserWord := 0

	for _, word := range words {
		wordsMap[word]++

		if wordsMap[word] > countOfMostUserWord {
			countOfMostUserWord = wordsMap[word]
			mostUserWord = word
		}
	}

	fmt.Println("summary:")
	fmt.Println("count of unique words:", len(wordsMap))
	fmt.Println("most used word:", mostUserWord, "\b,", countOfMostUserWord, "times")
}
