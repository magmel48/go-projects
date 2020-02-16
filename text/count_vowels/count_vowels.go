/**
Enter a string and the program counts the number of vowels in the text.
For added complexity have it report a sum of each vowel found.
*/
package count_vowels

import (
	"fmt"
	"github.com/magmel48/go-projects/text/pig_latin"
	"github.com/magmel48/go-projects/utils"
)

func CountVowels() {
	text := utils.ReadConsole()

	countVowels := 0
	reportingMap := make(map[string]int)

	for _, c := range text {
		casted := string(c)

		if pig_latin.IsVowel(casted) {
			countVowels++
			reportingMap[casted]++
		}
	}

	fmt.Println("vowels:", countVowels)
	for k, v := range reportingMap {
		fmt.Println("count of", k+":", v)
	}
}
